package config

type Config struct {
	Server ServerConfig `yaml:"server"`
	Core   CoreConfig   `yaml:"core"`
}

type ServerConfig struct {
	Port           int `yaml:"port"`
	MaxConnections int `yaml:"maxConnections"`
}

type CoreConfig struct {
	CleanerInterval uint64 `yaml:"cleanerInterval"`
}
