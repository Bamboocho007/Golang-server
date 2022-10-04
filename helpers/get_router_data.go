package helpers

import (
	"github.com/gin-gonic/gin"
)

func GetRouterData[T any](c *gin.Context, key string) (T, bool) {
	value, exists := c.Get(key)
	return value.(T), exists
}
