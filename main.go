package main

import (
	"auth/config"
	"fmt"
	"strconv"

	// "new-app/database"
	"auth/server"
)

func main() {
	cfg := config.GetConfig()

	// db := database.NewPostgresDatabase(&cfg)

	fmt.Println("the server is running on port " + strconv.Itoa(cfg.App.Port))
	server.NewEchoServer(&cfg).Start()
}
