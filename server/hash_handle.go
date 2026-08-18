package server

import memapv1 "github.com/memap-project/memap-proto/gen/memapv1/go"

func (s *Server) handleHGET(req *memapv1.Request) *memapv1.Response {
	hmap, err := s.manager.HGet(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okHashValue(hmap)
}

func (s *Server) handleHSET(req *memapv1.Request) *memapv1.Response {
	if err := s.manager.HSet(req.GetNamespace(), req.GetKey(), req.GetTtl()); err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleHDEL(req *memapv1.Request) *memapv1.Response {
	if err := s.manager.HDel(req.GetNamespace(), req.GetKey()); err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleHEXPIRE(req *memapv1.Request) *memapv1.Response {
	if err := s.manager.HExpire(req.GetNamespace(), req.GetKey(), req.GetTtl()); err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleHTTL(req *memapv1.Request) *memapv1.Response {
	ttl, err := s.manager.HTTL(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okIntValue(ttl)
}

func (s *Server) handleHEXIST(req *memapv1.Request) *memapv1.Response {
	exists, err := s.manager.HExists(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	if exists {
		return okEmpty()
	}
	return errEmpty()
}

func (s *Server) handleHLEN(req *memapv1.Request) *memapv1.Response {
	len, err := s.manager.HLen(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okIntValue(len)
}

func (s *Server) handleHKEYS(req *memapv1.Request) *memapv1.Response {
	keys, err := s.manager.HKeys(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okListValue(keys)
}

func (s *Server) handleHVALS(req *memapv1.Request) *memapv1.Response {
	values, err := s.manager.HValues(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okListValue(values)
}

func (s *Server) handleHFSET(req *memapv1.Request) *memapv1.Response {
	if err := s.manager.HFSet(req.GetNamespace(), req.GetKey(), req.GetField(), req.GetStringValue()); err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleHFGET(req *memapv1.Request) *memapv1.Response {
	value, err := s.manager.HFGet(req.GetNamespace(), req.GetKey(), req.GetField())
	if err != nil {
		return errResponse(err)
	}
	return okValue(value)
}

func (s *Server) handleHFDEL(req *memapv1.Request) *memapv1.Response {
	if err := s.manager.HFDel(req.GetNamespace(), req.GetKey(), req.GetField()); err != nil {
		return errResponse(err)
	}
	return okEmpty()
}
