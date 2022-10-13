package middlewares

import (
	"errors"
	"net/http"
	"strings"
	helpers "vasya_project/pkg/helpers"
	services "vasya_project/pkg/services"
	"vasya_project/pkg/utils"

	"github.com/gin-gonic/gin"
)

func CheckAuthorization(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"]
	val, exists := helpers.GetRouterData[utils.EnvConfig](c, "envConfig")

	if !exists {
		c.Abort()
		panic(errors.New("envConfig not exists"))
	}

	if authHeader == nil {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	_, err := services.Parse(strings.Split(strings.Join(authHeader, ""), " ")[1], val.Jwt_secret)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token",
		})
		c.Abort()
		return
	}

	c.Next()
}
