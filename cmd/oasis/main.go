package main

import (
	"fmt"
	"oasis/app"
	"oasis/cmd"
	"os"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Errorf("Fatal error config settings: %s \n", err)
		os.Exit(1)
	}
	app.RunServer()
}
