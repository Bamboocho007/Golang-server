package main

import (
	"vasya_project/pkg/router"
	"vasya_project/pkg/utils"
)

func main() {
	config, getConfirErr := utils.LoadConfig("development")

	if getConfirErr != nil {
		panic(getConfirErr)
	}

	err := router.InitRouter(config)
	println("some error %v", err)

}
