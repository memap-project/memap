package config

type Config struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Port           int `yaml:"port"`
	MaxConnections int `yaml:"maxConnections"`
}
