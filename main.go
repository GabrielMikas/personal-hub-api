package main

import (
	"fmt"

	"github.com/GabrielMikas/personal-hub-api/config"
	"github.com/GabrielMikas/personal-hub-api/router"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	router.Run()

}
