package config

type (
	// Config ...
	Config struct {
		Server   ServerConfig   `yaml:"server"`
		Database DatabaseConfig `yaml:"database"`
		Firebase FirebaseConfig `yaml:"firebase"`
	}

	// ServerConfig ...
	ServerConfig struct {
		Port string `yaml:"port"`
	}

	// DatabaseConfig ...
	DatabaseConfig struct {
		Master string `yaml:"master"`
	}
	// FirebaseConfig ...
	FirebaseConfig struct {
		ProjectID string `yaml:"projectID"`
	}
)
