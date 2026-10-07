package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/dmi3midd/protorw"
	"github.com/memap-project/memap-core/ns"
	memapv1 "github.com/memap-project/memap-proto/gen/memapv1/go"
	"github.com/memap-project/memap/config"
)

type Server struct {
	cfg        *config.ServerConfig
	semaphore  chan struct{}
	manager    *ns.NamespaceManager
	listener   net.Listener
	wg         sync.WaitGroup
	inShutdown atomic.Bool
	conns      sync.Map // tracks active net.Conn -> struct{}
}

func NewServer(
	cfg *config.ServerConfig,
	manager *ns.NamespaceManager,
) *Server {
	return &Server{
		cfg:       cfg,
		manager:   manager,
		semaphore: make(chan struct{}, cfg.MaxConnections),
	}
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.cfg.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	s.listener = listener
	slog.Info(
		"server started",
		slog.Int("port", s.cfg.Port),
		slog.Int("max_connections", s.cfg.MaxConnections),
		slog.Duration("idle_timeout", time.Duration(s.cfg.IdleTimeout)*time.Second),
	)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			if s.inShutdown.Load() || errors.Is(err, net.ErrClosed) {
				return nil
			}
			return fmt.Errorf("failed to accept: %w", err)
		}

		select {
		case s.semaphore <- struct{}{}:
			s.wg.Add(1)
			go func(c net.Conn) {
				defer func() {
					<-s.semaphore
					s.wg.Done()
				}()
				s.handleConnection(c)
			}(conn)
		default:
			// Виправляємо (А): Синхронна відмова з коротким дедлайном без створення зайвих горутин
			_ = conn.SetWriteDeadline(time.Now().Add(500 * time.Millisecond))
			resp := &memapv1.Response{
				Success: false,
				Error:   "server limit reached: too many connections",
			}
			_ = protorw.WriteMsg(conn, resp)
			_ = conn.Close()
		}
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.inShutdown.Store(true)
	var err error
	if s.listener != nil {
		err = s.listener.Close()
	}

	// Виправляємо (Г): Закриваємо всі активні з'єднання, перериваючи блокуючі ReadMsg
	s.conns.Range(func(key, _ any) bool {
		if c, ok := key.(net.Conn); ok {
			_ = c.Close()
		}
		return true
	})

	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	// Виправляємо (Б): Захист від панік, щоб збій одного запиту не клав увесь сервер
	defer func() {
		if r := recover(); r != nil {
			slog.Error("recovered from panic in connection handler",
				slog.Any("panic", r),
				slog.String("remote_addr", conn.RemoteAddr().String()),
			)
		}
		s.conns.Delete(conn)
		_ = conn.Close()
	}()

	// Виправляємо (Г): Реєструємо з'єднання в списку активних
	s.conns.Store(conn, struct{}{})

	for {
		if s.inShutdown.Load() {
			return
		}
		conn.SetDeadline(time.Now().Add(time.Duration(s.cfg.IdleTimeout) * time.Second))
		var req memapv1.Request
		err := protorw.ReadMsg(conn, &req)
		if err != nil {
			if errors.Is(err, io.EOF) || errors.Is(err, net.ErrClosed) {
				return
			}

			// Виправляємо (В): Не логуємо штатні таймаути неактивності як Error
			var netErr net.Error
			if errors.As(err, &netErr) && netErr.Timeout() {
				slog.Debug("connection closed due to idle timeout",
					slog.String("remote_addr", conn.RemoteAddr().String()),
				)
				return
			}

			// Якщо сервер завершує роботу, розрив сокетів є очікуваним
			if s.inShutdown.Load() {
				return
			}

			slog.Error("failed to read", slog.String("error", err.Error()))
			return
		}

		resp := s.processRequest(&req)

		if err := protorw.WriteMsg(conn, resp); err != nil {
			if !s.inShutdown.Load() && !errors.Is(err, net.ErrClosed) {
				slog.Error("failed to write", slog.String("error", err.Error()))
			}
			return
		}
	}
}
