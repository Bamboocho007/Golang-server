package middlewares

import (
	"errors"
	"net/http"
	"strings"
	helpers "vasya_project/helpers"
	services "vasya_project/services"
	"vasya_project/utils"

	"github.com/gin-gonic/gin"
)

func CheckRole(accept []int) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header["Authorization"]

		val, exists := helpers.GetRouterData[utils.EnvConfig](c, "envConfig")

		if !exists {
			c.Abort()
			panic(errors.New("envConfig not exists"))
		}

		if authHeader == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "No auth hrader",
			})
			c.Abort()
			return
		}

		claims, err := services.Parse(strings.Split(strings.Join(authHeader, ""), " ")[1], val.Jwt_secret)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			c.Abort()
			return
		}

		hasPermissions := false

		for _, item := range accept {
			println("item", item)
			println("claims.Role", claims.Role)
			if item == claims.Role {
				hasPermissions = true
				c.Next()
				break
			}
		}

		if !hasPermissions {
			c.JSON(http.StatusForbidden, gin.H{})
			c.Abort()
			return
		}
	}
}
