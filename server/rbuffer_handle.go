package server

import memapv1 "github.com/memap-project/memap-proto/gen/memapv1/go"

func (s *Server) handleBINIT(req *memapv1.Request) *memapv1.Response {
	err := s.manager.BInit(req.GetNamespace(), req.GetKey(), req.GetLimit(), req.GetTtl())
	if err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleBPUSH(req *memapv1.Request) *memapv1.Response {
	err := s.manager.BPush(req.GetNamespace(), req.GetKey(), req.GetStringValue())
	if err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleBPOP(req *memapv1.Request) *memapv1.Response {
	val, err := s.manager.BPop(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okValue(val)
}

func (s *Server) handleBAT(req *memapv1.Request) *memapv1.Response {
	val, err := s.manager.BAt(req.GetNamespace(), req.GetKey(), req.GetIntValue())
	if err != nil {
		return errResponse(err)
	}
	return okValue(val)
}

func (s *Server) handleBSLICE(req *memapv1.Request) *memapv1.Response {
	val, err := s.manager.BSlice(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okListValue(val)
}

func (s *Server) handleBPEEK(req *memapv1.Request) *memapv1.Response {
	val, err := s.manager.BPeek(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okValue(val)
}

func (s *Server) handleBBACK(req *memapv1.Request) *memapv1.Response {
	val, err := s.manager.BBack(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okValue(val)
}

func (s *Server) handleBCAP(req *memapv1.Request) *memapv1.Response {
	val, err := s.manager.BCap(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okIntValue(val)
}

func (s *Server) handleBLEN(req *memapv1.Request) *memapv1.Response {
	val, err := s.manager.BLen(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okIntValue(val)
}

func (s *Server) handleBRESET(req *memapv1.Request) *memapv1.Response {
	err := s.manager.BReset(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleBDEL(req *memapv1.Request) *memapv1.Response {
	err := s.manager.BDel(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleBEXPIRE(req *memapv1.Request) *memapv1.Response {
	err := s.manager.BExpire(req.GetNamespace(), req.GetKey(), req.GetTtl())
	if err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleBTTL(req *memapv1.Request) *memapv1.Response {
	val, err := s.manager.BTTL(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okIntValue(val)
}
