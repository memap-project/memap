package server

import memapv1 "github.com/memap-project/memap-proto/gen/memapv1/go"

func (s *Server) handleSADD(req *memapv1.Request) *memapv1.Response {
	err := s.manager.SAdd(req.GetNamespace(), req.GetKey(), req.GetStringValue(), req.GetTtl())
	if err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleSREMOVE(req *memapv1.Request) *memapv1.Response {
	err := s.manager.SRemove(req.GetNamespace(), req.GetKey(), req.GetStringValue())
	if err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleSISMEMBER(req *memapv1.Request) *memapv1.Response {
	isMember, err := s.manager.SIsMember(req.GetNamespace(), req.GetKey(), req.GetStringValue())
	if err != nil {
		return errResponse(err)
	}

	if isMember {
		return okIntValue(1)
	}
	return okIntValue(0)
}

func (s *Server) handleSCARD(req *memapv1.Request) *memapv1.Response {
	val, err := s.manager.SCard(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okIntValue(val)
}

func (s *Server) handleSMEMBERS(req *memapv1.Request) *memapv1.Response {
	val, err := s.manager.SMembers(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okListValue(val)
}

func (s *Server) handleSEXPIRE(req *memapv1.Request) *memapv1.Response {
	err := s.manager.SExpire(req.GetNamespace(), req.GetKey(), req.GetTtl())
	if err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleSTTL(req *memapv1.Request) *memapv1.Response {
	val, err := s.manager.STTL(req.GetNamespace(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okIntValue(val)
}
