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
	CleanerInterval int             `yaml:"cleanerInterval"`
	Namespace       NamespaceConfig `yaml:"namespace"`
}

type NamespaceConfig struct {
	ShardCounts ShardCounts `yaml:"shardCounts"`
}

type ShardCounts struct {
	Shmap     uint8 `yaml:"shmap"`
	Shhash    uint8 `yaml:"shhash"`
	Shcounter uint8 `yaml:"shcounter"`
	Shrbuffer uint8 `yaml:"shrbuffer"`
}
