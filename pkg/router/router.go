package router

import (
	middlewares "vasya_project/pkg/middlewares"
	"vasya_project/pkg/utils"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engin  *gin.Engine
	config utils.EnvConfig
}

func InitRouter(config utils.EnvConfig) error {
	r := Router{engin: gin.Default(), config: config}
	r.engin.Use(middlewares.EnvConfigExtention(config))
	r.addUsers()
	r.addAuth()
	err := r.engin.Run()
	return err
}
