package config

import (
	coreconfig "github.com/memap-project/memap-core/config"
)

type Config struct {
	Server ServerConfig      `yaml:"server"`
	Core   coreconfig.Config `yaml:"core"`
}

type ServerConfig struct {
	Port           int `yaml:"port"`
	MaxConnections int `yaml:"maxConnections"`
}
