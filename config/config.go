package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const (
	SetMaxOpenConns = 5
	SetMaxIdleConns = 3
)

var (
	DB *gorm.DB
)

type Config struct {
	Server Server
	MySQL  MySQLConfig
}

type Server struct {
	Port          string
	LogError      string `toml:"error_log"`
	LogAccess     string `toml:"access_log"`
	LogAccessPath string `toml:"access_log_path"`
	Bind          string
}

type MySQLConfig struct {
	Host        string
	Port        string
	User        string
	Password    string
	Database    string
	MaxOpenConn int
	MaxIdleConn int
}

func NewConfig() *Config {
	return &Config{
		Server: Server{
			Port:          viper.GetString("server.port"),
			LogError:      viper.GetString("server.error_log"),
			LogAccess:     viper.GetString("server.access_log"),
			LogAccessPath: viper.GetString("server.access_log_path"),
			Bind:          viper.GetString("server.bind"),
		},
		MySQL: MySQLConfig{
			User:        viper.GetString("mysql.user"),
			Password:    viper.GetString("mysql.password"),
			Host:        viper.GetString("mysql.host"),
			Port:        viper.GetString("mysql.port"),
			Database:    viper.GetString("mysql.database"),
			MaxOpenConn: SetMaxOpenConns,
			MaxIdleConn: SetMaxIdleConns,
		},
	}
}
