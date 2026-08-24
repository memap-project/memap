package server

import (
	"fmt"
	"io"
	"net"

	"github.com/dmi3midd/protorw"
	"github.com/memap-project/memap-core/ns"
	"github.com/memap-project/memap/config"

	memapv1 "github.com/memap-project/memap-proto/gen/memapv1/go"
)

type Server struct {
	cfg       *config.ServerConfig
	semaphore chan struct{}
	manager   *ns.NamespaceManager
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
	defer listener.Close()
	fmt.Println("server started")
	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("failed to accept: %w", err)
		}
		select {
		case s.semaphore <- struct{}{}:
			go func() {
				defer func() {
					<-s.semaphore
				}()
				s.handleConnection(conn)
			}()
		default:
			go func(c net.Conn) {
				defer c.Close()
				resp := &memapv1.Response{
					Success: false,
					Error:   "server limit reached: too many connections",
				}
				_ = protorw.WriteMsg(c, resp)
			}(conn)
		}
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		var req memapv1.Request
		err := protorw.ReadMsg(conn, &req)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Printf("failed to read: %v\n", err)
			return
		}
		fmt.Println(req.String())

		resp := s.processRequest(&req)

		if err := protorw.WriteMsg(conn, resp); err != nil {
			fmt.Printf("failed to write: %v\n", err)
			return
		}
	}
}
