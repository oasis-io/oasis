package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"oasis/pkg/log"
)

const (
	VERSION              = "0.1.0"     // Oasis version
	DefaultAdminUsername = "admin"     // 超级管理员用户名
	DefaultAdminPassword = "Oasis2022" // 超级管理员密码

	SetMaxOpenConns = 5
	SetMaxIdleConns = 3

	ServerBind = "server.bind"
)

var (
	DB *gorm.DB
)

type OasisConfig struct {
	Server Server
	MySQL  MySQLConfig
}

type Server struct {
	Port          string `toml:"port"`
	LogError      string `toml:"error_log"`
	LogAccess     string `toml:"access_log"`
	LogAccessPath string `toml:"access_log_path"`
	Bind          string `toml:"bind"`
}

type MySQLConfig struct {
	Host        string `toml:"host"`
	Port        string `toml:"port"`
	User        string `toml:"user"`
	Password    string `toml:"password"`
	Database    string `toml:"database"`
	MaxOpenConn int
	MaxIdleConn int
}

func NewOasisConfig() *OasisConfig {
	return &OasisConfig{
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

func InitConfig(cfgFile string) error {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("toml")
		viper.SetConfigName("oasis")

		// server default value
		viper.SetDefault("server.bind", "127.0.0.1")
		viper.SetDefault("server.port", "9590")
		viper.SetDefault("server.error_log", "./oasis.log")
		viper.SetDefault("server.access_log", "off")
		viper.SetDefault("server.access_log_path", "./access.log")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	log.Info("Initialized config file", zap.String("file", viper.ConfigFileUsed()))

	return nil
}
