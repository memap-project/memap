package server

import memapv1 "github.com/memap-project/memap-proto/gen/memapv1/go"

func (s *Server) handleSET(req *memapv1.Request) *memapv1.Response {
	err := s.manager.Set(req.GetNamespaceName(), req.GetKey(), req.GetValue(), req.GetTtlSeconds())
	if err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleGET(req *memapv1.Request) *memapv1.Response {
	v, err := s.manager.Get(req.GetNamespaceName(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okValue(v)
}

func (s *Server) handleDEL(req *memapv1.Request) *memapv1.Response {
	if err := s.manager.Delete(req.GetNamespaceName(), req.GetKey()); err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleEXPIRE(req *memapv1.Request) *memapv1.Response {
	if err := s.manager.Expire(req.GetNamespaceName(), req.GetKey(), req.GetTtlSeconds()); err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleTTL(req *memapv1.Request) *memapv1.Response {
	ttl, err := s.manager.TTL(req.GetNamespaceName(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okIntValue(ttl)
}
