package main

import (
	"loggerservice/pkg/config"
	"loggerservice/pkg/data"
)

func main() {

	// Initialize environment variables
	config.InitEnvConfigs()

	client := data.InitDB()
	defer data.CloseDB(client)

}
