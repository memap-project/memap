package config

import (
	coreconfig "github.com/memap-project/memap-core/config"
)

type Config struct {
	Server ServerConfig      `yaml:"server"`
	Logger LoggerConfig      `yaml:"logger"`
	Core   coreconfig.Config `yaml:"core"`
}

type LoggerConfig struct {
	LogPath string `yaml:"logPath"`
}

type ServerConfig struct {
	Port           int `yaml:"port"`
	MaxConnections int `yaml:"maxConnections"`
	IdleTimeout    int `yaml:"idleTimeout"`
}
