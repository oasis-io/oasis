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
	Port string
	Log  string `toml:"log_error"`
	Bind string
}

type MySQLConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	Database     string
	MaxOpenConns int
	MaxIdleConns int
}

func NewConfig() *Config {
	return &Config{
		Server: Server{
			Port: viper.GetString("server.port"),
			Log:  viper.GetString("server.log_error"),
			Bind: viper.GetString("server.bind"),
		},
		MySQL: MySQLConfig{
			User:         viper.GetString("mysql.user"),
			Password:     viper.GetString("mysql.password"),
			Host:         viper.GetString("mysql.host"),
			Port:         viper.GetString("mysql.port"),
			Database:     viper.GetString("mysql.database"),
			MaxOpenConns: SetMaxOpenConns,
			MaxIdleConns: SetMaxIdleConns,
		},
	}
}
