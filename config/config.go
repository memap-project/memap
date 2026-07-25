package config

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port           int
	MaxConnections int
}
