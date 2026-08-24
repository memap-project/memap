package config

import (
	"errors"
)

var (
	ErrInvalidPort = errors.New("server.port must be between 1 and 65535")
)

// validatePort returns an error if the port is not between 1 and 65535.
func (c *Config) validatePort() error {
	if c.Server.Port <= 0 || c.Server.Port > 65535 {
		return ErrInvalidPort
	}
	return nil
}

func (c *Config) Validate() error {
	if err := c.validatePort(); err != nil {
		return err
	}
	if err := c.Core.Validate(); err != nil {
		return err
	}
	return nil
}
