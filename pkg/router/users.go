package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r Router) addUsers() {
	group := r.engin.Group("users")
	// group.Use(middlewares.CheckRole([]int{0, 1}))
	group.GET(":id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"userId": id,
		})
	})
}
