package config

import (
	"errors"
)

var (
	ErrInvalidPort           = errors.New("server.port must be between 1 and 65535")
	ErrInvalidIdleTimeout    = errors.New("idle timeout must be greater than 0")
	ErrInvalidMaxConnections = errors.New("max connections must be greater than 0")
)

// validatePort returns an error if the port is not between 1 and 65535.
func (c *Config) validatePort() error {
	if c.Server.Port <= 0 || c.Server.Port > 65535 {
		return ErrInvalidPort
	}
	return nil
}

func (c *Config) validateMaxConnections() error {
	if c.Server.MaxConnections <= 0 {
		return ErrInvalidMaxConnections
	}
	return nil
}

func (c *Config) validateIdleTimeout() error {
	if c.Server.IdleTimeout <= 0 {
		return ErrInvalidIdleTimeout
	}
	return nil
}

func (c *Config) Validate() error {
	if err := c.validatePort(); err != nil {
		return err
	}
	if err := c.validateIdleTimeout(); err != nil {
		return err
	}
	if err := c.validateMaxConnections(); err != nil {
		return err
	}
	if err := c.Core.Validate(); err != nil {
		return err
	}
	return nil
}
