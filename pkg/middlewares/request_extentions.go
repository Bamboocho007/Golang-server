package middlewares

import (
	"vasya_project/pkg/utils"

	"github.com/gin-gonic/gin"
)

func EnvConfigExtention(config utils.EnvConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("envConfig", config)
		c.Next()
	}
}
