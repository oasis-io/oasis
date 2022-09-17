package cmd

import (
	"oasis/api"
	"oasis/config"
	"oasis/log"
	"oasis/model"
)

func Run(config *config.Config) error {
	// 初始化日志库
	log.InitLogger(config.Server.LogPath)
	defer log.ExitLogger()
	log.Logger().Infoln("starting oasis server")

	// 初始化表结构
	model.Migrate(config)

	// 启动路由
	api.Routers()

	return nil
}
