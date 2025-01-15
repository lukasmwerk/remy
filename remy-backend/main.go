package main

import (
	"fmt"
	"log"

	"github.com/lukasmwerk/remy/remy-backend/cmd"
	"github.com/lukasmwerk/remy/remy-backend/config"
)

func main() {
	config, err := config.LoadConfiguration("../data", "config", "yaml")
	if err != nil {
		log.Fatal("could not load configuration file: %v", err)
	}
	fmt.Println(config.Environment)
	cmd.Execute()
}
