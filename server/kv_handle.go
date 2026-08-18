package server

import memapv1 "github.com/memap-project/memap-proto/gen/memapv1/go"

func (s *Server) handleSET(req *memapv1.Request) *memapv1.Response {
	err := s.manager.Set(req.GetNamespace(), req.GetKey(), req.GetStringValue(), req.GetTtl())
	if err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleGET(req *memapv1.Request) *memapv1.Response {
	v, err := s.manager.Get(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okValue(v)
}

func (s *Server) handleDEL(req *memapv1.Request) *memapv1.Response {
	if err := s.manager.Del(req.GetNamespace(), req.GetKey()); err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleEXPIRE(req *memapv1.Request) *memapv1.Response {
	if err := s.manager.Expire(req.GetNamespace(), req.GetKey(), req.GetTtl()); err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleTTL(req *memapv1.Request) *memapv1.Response {
	ttl, err := s.manager.TTL(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okIntValue(ttl)
}
