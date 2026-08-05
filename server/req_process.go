package server

import (
	memapv1 "github.com/memap-project/memap-proto/gen/memapv1/go"
)

type commandHandler func(s *Server, req *memapv1.Request) *memapv1.Response

var commandHandlers = map[memapv1.CommandType]commandHandler{
	memapv1.CommandType_CREATE_NS: (*Server).handleCREATE_NS,
	memapv1.CommandType_DELETE_NS: (*Server).handleDELETE_NS,
	memapv1.CommandType_SET:       (*Server).handleSET,
	memapv1.CommandType_GET:       (*Server).handleGET,
	memapv1.CommandType_DEL:       (*Server).handleDEL,
	memapv1.CommandType_EXPIRE:    (*Server).handleEXPIRE,
	memapv1.CommandType_HGET:      (*Server).handleHGET,
	memapv1.CommandType_HSET:      (*Server).handleHSET,
	memapv1.CommandType_HDEL:      (*Server).handleHDEL,
	memapv1.CommandType_HEXPIRE:   (*Server).handleHEXPIRE,
	memapv1.CommandType_HFSET:     (*Server).handleHFSET,
	memapv1.CommandType_HFGET:     (*Server).handleHFGET,
	memapv1.CommandType_HFDEL:     (*Server).handleHFDEL,
	memapv1.CommandType_PING: func(s *Server, req *memapv1.Request) *memapv1.Response {
		return okValue("PONG")
	},
}

func (s *Server) processRequest(req *memapv1.Request) *memapv1.Response {
	handler, ok := commandHandlers[req.Type]
	if !ok {
		return &memapv1.Response{
			Success:      false,
			ErrorMessage: "unknown command",
		}
	}
	return handler(s, req)
}

func okValue(v string) *memapv1.Response {
	return &memapv1.Response{Success: true, Value: v}
}

func okHash(h map[string]string) *memapv1.Response {
	return &memapv1.Response{Success: true, HashValue: h}
}

func okEmpty() *memapv1.Response {
	return &memapv1.Response{Success: true}
}

func errResponse(err error) *memapv1.Response {
	return &memapv1.Response{Success: false, ErrorMessage: err.Error()}
}
