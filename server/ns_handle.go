package server

import memapv1 "github.com/memap-project/memap-proto/gen/memapv1/go"

func (s *Server) handleCREATE_NS(req *memapv1.Request) *memapv1.Response {
	if err := s.manager.CreateNs(req.GetNamespaceName()); err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleDELETE_NS(req *memapv1.Request) *memapv1.Response {
	if err := s.manager.DeleteNs(req.GetNamespaceName()); err != nil {
		return errResponse(err)
	}
	return okEmpty()
}
