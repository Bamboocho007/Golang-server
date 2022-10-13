package router

import (
	"net/http"
	services "vasya_project/pkg/services"

	"github.com/gin-gonic/gin"
)

func (r Router) addAuth() {
	group := r.engin.Group("auth")

	group.POST("/login", func(c *gin.Context) {
		res_data := services.Login(c)
		c.JSON(http.StatusOK, res_data)
	})
}
