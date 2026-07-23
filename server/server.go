package server

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/dmi3midd/memap/config"
)

type Server struct {
	cfg *config.ServerConfig
}

func NewServer(cfg *config.ServerConfig) *Server {
	return &Server{cfg: cfg}
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
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Read error: %v", err)
			}
			return
		}
		fmt.Println(bytes)
		fmt.Println(string(bytes))
		_, err = conn.Write([]byte(fmt.Sprintf("response: %s", string(bytes))))
		if err != nil {
			fmt.Printf("Write error: %v", err)
		}
	}
}
