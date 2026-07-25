package server

import (
	"time"

	memapv1 "github.com/memap-project/memap-proto/gen/memapv1/go"
)

func (s *Server) processRequest(req *memapv1.Request) *memapv1.Response {
	command := req.Type
	switch command {
	// SET
	case memapv1.CommandType_SET:
		ns := req.GetNamespaceName()
		key := req.GetKey()
		value := req.GetValue()
		ttl := req.GetTtlSeconds()

		err := s.manager.Set(ns, key, value, time.Duration(ttl))
		if err != nil {
			return &memapv1.Response{
				Success:      false,
				Value:        "",
				ErrorMessage: err.Error(),
			}
		}
		return &memapv1.Response{
			Success:      true,
			ErrorMessage: "",
		}

	// GET
	case memapv1.CommandType_GET:
		ns := req.GetNamespaceName()
		key := req.GetKey()

		value, err := s.manager.Get(ns, key)
		if err != nil {
			return &memapv1.Response{
				Success:      false,
				Value:        "",
				ErrorMessage: err.Error(),
			}
		}
		return &memapv1.Response{
			Success:      true,
			Value:        value.Value,
			ErrorMessage: "",
		}

	// DEL
	case memapv1.CommandType_DEL:
		ns := req.GetNamespaceName()
		key := req.GetKey()

		err := s.manager.Delete(ns, key)
		if err != nil {
			return &memapv1.Response{
				Success:      false,
				ErrorMessage: err.Error(),
			}
		}
		return &memapv1.Response{
			Success:      true,
			ErrorMessage: "",
		}

	// CREATE_NS
	case memapv1.CommandType_CREATE_NS:
		ns := req.GetNamespaceName()

		err := s.manager.CreateNs(ns)
		if err != nil {
			return &memapv1.Response{
				Success:      false,
				ErrorMessage: err.Error(),
			}
		}
		return &memapv1.Response{
			Success:      true,
			ErrorMessage: "",
		}

	// DELETE_NS
	case memapv1.CommandType_DELETE_NS:
		ns := req.GetNamespaceName()

		err := s.manager.DeleteNs(ns)
		if err != nil {
			return &memapv1.Response{
				Success:      false,
				ErrorMessage: err.Error(),
			}
		}
		return &memapv1.Response{
			Success:      true,
			ErrorMessage: "",
		}

	// PING
	case memapv1.CommandType_PING:
		return &memapv1.Response{
			Success:      true,
			Value:        "PONG",
			ErrorMessage: "",
		}

	}
	// UNKNOWN
	return &memapv1.Response{
		Success:      false,
		ErrorMessage: "unknown command",
	}
}
