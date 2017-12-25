package baseController

import (
	"github.com/gin-gonic/gin"
)

func (this *Base) Get(key string, defaultVal interface{}, c *gin.Context) interface{} {
	value, exist := c.Get(key)

	if !exist {
		return defaultVal
	}
	return value
}
