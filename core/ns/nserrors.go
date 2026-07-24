package ns

import "errors"

var (
	ErrNamespaceAlreadyExists = errors.New("namespace already exists")
	ErrNamespaceNotFound      = errors.New("namespace not found")
	ErrKeyNotFound            = errors.New("key not found")
	ErrExpiredKey             = errors.New("expired key")
	ErrInvalidNamespace       = errors.New("invalid namespace")
)
