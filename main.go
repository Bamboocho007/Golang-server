package main

import (
	"vasya_project/router"
	"vasya_project/utils"
)

func main() {
	config, getConfirErr := utils.LoadConfig("development")

	if getConfirErr != nil {
		panic(getConfirErr)
	}

	err := router.InitRouter(config)
	println("some error %v", err)

}
