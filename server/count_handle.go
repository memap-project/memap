package server

import memapv1 "github.com/memap-project/memap-proto/gen/memapv1/go"

func (s *Server) handleCINIT(req *memapv1.Request) *memapv1.Response {
	err := s.manager.Init(req.GetNamespace(), req.GetKey(), req.GetLimit(), req.GetTtl())
	if err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleSLIMIT(req *memapv1.Request) *memapv1.Response {
	err := s.manager.SLimit(req.GetNamespace(), req.GetKey(), req.GetLimit())
	if err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleGLIMIT(req *memapv1.Request) *memapv1.Response {
	limit, err := s.manager.GLimit(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okIntValue(limit)
}

func (s *Server) handleCGET(req *memapv1.Request) *memapv1.Response {
	count, err := s.manager.CGet(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okIntValue(count)
}

func (s *Server) handleCDEL(req *memapv1.Request) *memapv1.Response {
	err := s.manager.Del(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleCEXPIRE(req *memapv1.Request) *memapv1.Response {
	err := s.manager.CExpire(req.GetNamespace(), req.GetKey(), req.GetTtl())
	if err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleCTTL(req *memapv1.Request) *memapv1.Response {
	ttl, err := s.manager.CTTL(req.GetNamespace(), req.GetKey())
	if err != nil {
		resp := errResponse(err)
		resp.IntValue = ttl
		return resp
	}
	return okIntValue(ttl)
}

func (s *Server) handleINCRBY(req *memapv1.Request) *memapv1.Response {
	count, err := s.manager.IncrBy(req.GetNamespace(), req.GetKey(), req.GetIntValue())
	if err != nil {
		return errResponse(err)
	}
	return okIntValue(count)
}

func (s *Server) handleDECRBY(req *memapv1.Request) *memapv1.Response {
	count, err := s.manager.DecrBy(req.GetNamespace(), req.GetKey(), req.GetIntValue())
	if err != nil {
		return errResponse(err)
	}
	return okIntValue(count)
}
