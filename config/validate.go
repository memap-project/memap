package config

import (
	"errors"
)

var (
	ErrInvalidShardCount      = errors.New("shard count must be a power of two")
	ErrInvalidPort            = errors.New("server.port must be between 1 and 65535")
	ErrInvalidCleanerInterval = errors.New("core.cleanerInterval must be greater than 0")
)

func isPowerOfTwo(n uint8) bool {
	return n > 0 && (n&(n-1)) == 0
}

// validatePort returns an error if the port is not between 1 and 65535.
func (c *Config) validatePort() error {
	if c.Server.Port <= 0 || c.Server.Port > 65535 {
		return ErrInvalidPort
	}
	return nil
}

// validateCleanerInterval returns an error if the cleaner interval is not greater than 0.
func (c *Config) validateCleanerInterval() error {
	if c.Core.CleanerInterval == 0 {
		return ErrInvalidCleanerInterval
	}
	return nil
}

// validateShardCount returns an error if the shard count is not a power of two.
func (c *Config) validateShardCount() error {
	if !isPowerOfTwo(c.Core.Namespace.ShardCounts.Shmap) {
		return ErrInvalidShardCount
	}
	if !isPowerOfTwo(c.Core.Namespace.ShardCounts.Shhash) {
		return ErrInvalidShardCount
	}
	if !isPowerOfTwo(c.Core.Namespace.ShardCounts.Shcounter) {
		return ErrInvalidShardCount
	}
	return nil
}

func (c *Config) Validate() error {
	if err := c.validatePort(); err != nil {
		return err
	}
	if err := c.validateCleanerInterval(); err != nil {
		return err
	}
	if err := c.validateShardCount(); err != nil {
		return err
	}
	return nil
}
