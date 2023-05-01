package app

import (
	"oasis/config"
	db2 "oasis/db"
)

var err error

func RunServer() {
	// conn mysql
	config.DB, err = db2.OpenOasis()
	if err != nil {
		panic(err)
	}

	startHttp()
}

func startHttp() {
	HttpRequests()
}
