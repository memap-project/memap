package server

import memapv1 "github.com/memap-project/memap-proto/gen/memapv1/go"

func (s *Server) handleHGET(req *memapv1.Request) *memapv1.Response {
	hmap, err := s.manager.HGet(req.GetNamespaceName(), req.GetKey())
	if err != nil {
		return errResponse(err)
	}
	return okHash(hmap)
}

func (s *Server) handleHSET(req *memapv1.Request) *memapv1.Response {
	if err := s.manager.HSet(req.GetNamespaceName(), req.GetKey(), req.GetTtlSeconds()); err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleHDEL(req *memapv1.Request) *memapv1.Response {
	if err := s.manager.HDelete(req.GetNamespaceName(), req.GetKey()); err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleHFSET(req *memapv1.Request) *memapv1.Response {
	if err := s.manager.HFSet(req.GetNamespaceName(), req.GetKey(), req.GetField(), req.GetValue(), req.GetTtlSeconds()); err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleHFGET(req *memapv1.Request) *memapv1.Response {
	value, err := s.manager.HFGet(req.GetNamespaceName(), req.GetKey(), req.GetField())
	if err != nil {
		return errResponse(err)
	}
	return okValue(value)
}

func (s *Server) handleHFDEL(req *memapv1.Request) *memapv1.Response {
	if err := s.manager.HFDelete(req.GetNamespaceName(), req.GetKey(), req.GetField()); err != nil {
		return errResponse(err)
	}
	return okEmpty()
}
