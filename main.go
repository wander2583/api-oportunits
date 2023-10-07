package main

import (
	"fmt"

	"github.com/wander2583/api-oportunits/config"
	"github.com/wander2583/api-oportunits/router"
)

func main() {

	// Initialize Configs
	err := config.Init()
	if err != nil {
		panic(err)
		fmt.Println(err)
		return
	}

	// Initialize Router
	router.Initialize()
}
