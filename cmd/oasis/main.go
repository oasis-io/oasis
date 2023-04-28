package main

import (
	"oasis/app"
	"oasis/cmd"
)

func main() {
	cmd.Execute()
	app.RunServer()
}
