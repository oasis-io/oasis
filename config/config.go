package config

type Config struct {
	Server   Server   `yaml:"global"`
	DBConfig MyConfig `yaml:"db_config"`
}

type Server struct {
	LogPath string `yaml:"log_path"`
}

type MyConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
