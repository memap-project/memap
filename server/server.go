package server

import (
	"fmt"
	"io"
	"net"

	"github.com/dmi3midd/memap/config"
	"github.com/dmi3midd/memap/core/ns"
	"github.com/dmi3midd/protorw"

	memapv1 "github.com/dmi3midd/memap/proto/gen/memapv1/go"
)

// TODO: Implement Worker Pool pattern for server
type Server struct {
	cfg     *config.ServerConfig
	manager *ns.NamespaceManager
	pool    *Pool
}

func NewServer(
	cfg *config.ServerConfig,
) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Start() error {
	s.pool = NewPool(s.handleConnection, 100)
	s.pool.Start(10)

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

		s.pool.tasks <- conn
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
			fmt.Printf("Read error: %v\n", err)
			return
		}
		fmt.Println(req.String())

		resp := s.processRequest(&req)

		if err := protorw.WriteMsg(conn, resp); err != nil {
			fmt.Printf("Write error: %v\n", err)
			return
		}
	}
}
