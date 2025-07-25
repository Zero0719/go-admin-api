package main

import (
	"fmt"
	"go-admin-api/app/config"
)

func main() {
	fmt.Println("Hello, World!")
	config.Init()
	fmt.Println(config.Cfg.Server.Port)
}
