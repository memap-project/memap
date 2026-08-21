package server

import (
	memapv1 "github.com/memap-project/memap-proto/gen/memapv1/go"
)

type commandHandler func(s *Server, req *memapv1.Request) *memapv1.Response

var commandHandlers = map[memapv1.CommandType]commandHandler{
	memapv1.CommandType_CREATE: (*Server).handleCREATE,
	memapv1.CommandType_DROP:   (*Server).handleDROP,
	memapv1.CommandType_ERASE:  (*Server).handleERASE,
	memapv1.CommandType_FLUSH:  (*Server).handleFLUSH,

	memapv1.CommandType_SET:    (*Server).handleSET,
	memapv1.CommandType_GET:    (*Server).handleGET,
	memapv1.CommandType_DEL:    (*Server).handleDEL,
	memapv1.CommandType_EXPIRE: (*Server).handleEXPIRE,
	memapv1.CommandType_TTL:    (*Server).handleTTL,

	memapv1.CommandType_HGET:    (*Server).handleHGET,
	memapv1.CommandType_HSET:    (*Server).handleHSET,
	memapv1.CommandType_HDEL:    (*Server).handleHDEL,
	memapv1.CommandType_HEXPIRE: (*Server).handleHEXPIRE,
	memapv1.CommandType_HTTL:    (*Server).handleHTTL,
	memapv1.CommandType_HEXIST:  (*Server).handleHEXIST,
	memapv1.CommandType_HLEN:    (*Server).handleHLEN,
	memapv1.CommandType_HKEYS:   (*Server).handleHKEYS,
	memapv1.CommandType_HVALS:   (*Server).handleHVALS,
	memapv1.CommandType_HFSET:   (*Server).handleHFSET,
	memapv1.CommandType_HFGET:   (*Server).handleHFGET,
	memapv1.CommandType_HFDEL:   (*Server).handleHFDEL,

	memapv1.CommandType_CINIT:   (*Server).handleCINIT,
	memapv1.CommandType_CSLIMIT: (*Server).handleCSLIMIT,
	memapv1.CommandType_CGLIMIT: (*Server).handleCGLIMIT,
	memapv1.CommandType_CGET:    (*Server).handleCGET,
	memapv1.CommandType_CDEL:    (*Server).handleCDEL,
	memapv1.CommandType_CEXPIRE: (*Server).handleCEXPIRE,
	memapv1.CommandType_CTTL:    (*Server).handleCTTL,
	memapv1.CommandType_CINCRBY: (*Server).handleCINCRBY,
	memapv1.CommandType_CDECRBY: (*Server).handleCDECRBY,

	memapv1.CommandType_PING: func(s *Server, req *memapv1.Request) *memapv1.Response {
		return okValue("PONG")
	},
}

func (s *Server) processRequest(req *memapv1.Request) *memapv1.Response {
	handler, ok := commandHandlers[req.Command]
	if !ok {
		return &memapv1.Response{
			Success: false,
			Error:   "unknown command",
		}
	}
	return handler(s, req)
}

func okValue(v string) *memapv1.Response {
	return &memapv1.Response{Success: true, StringValue: v}
}

func okIntValue(v int64) *memapv1.Response {
	return &memapv1.Response{Success: true, IntValue: v}
}

func okHashValue(h map[string]string) *memapv1.Response {
	return &memapv1.Response{Success: true, MapValue: h}
}

func okListValue(l []string) *memapv1.Response {
	return &memapv1.Response{Success: true, SliceValue: l}
}

func okEmpty() *memapv1.Response {
	return &memapv1.Response{Success: true}
}

func errEmpty() *memapv1.Response {
	return &memapv1.Response{Success: false}
}

func errResponse(err error) *memapv1.Response {
	return &memapv1.Response{Success: false, Error: err.Error()}
}
