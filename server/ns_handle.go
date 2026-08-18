package server

import memapv1 "github.com/memap-project/memap-proto/gen/memapv1/go"

func (s *Server) handleCREATE(req *memapv1.Request) *memapv1.Response {
	if err := s.manager.Create(req.GetNamespace()); err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleDROP(req *memapv1.Request) *memapv1.Response {
	if err := s.manager.Drop(req.GetNamespace()); err != nil {
		return errResponse(err)
	}
	return okEmpty()
}

func (s *Server) handleERASE(_ *memapv1.Request) *memapv1.Response {
	s.manager.Erase()
	return okEmpty()
}

func (s *Server) handleFLUSH(_ *memapv1.Request) *memapv1.Response {
	s.manager.Flush()
	return okEmpty()
}
