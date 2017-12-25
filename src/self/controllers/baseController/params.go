package baseController

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	ers "self/commons/g"
)

func (this *Base) MustString(key string, c *gin.Context) string {
	return this.MustGetStrParam(key, c)
}

func (this *Base) String(key, defaultVal string, c *gin.Context) string {
	ret := this.getStrParam(key, c)
	if ret == "" {
		return defaultVal
	}
	return ret
}

func (this *Base) MustInt(key string, c *gin.Context) int {
	param := this.MustGetStrParam(key, c)
	ret, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		errInfo := fmt.Sprintf("参数[%s]非int型，请检查", key)
		panic(ers.ParamError(errInfo))
	}
	return int(ret)
}

func (this *Base) Int(key string, defaultVal int, c *gin.Context) int {
	param := this.getStrParam(key, c)
	if param == "" {
		return defaultVal
	}
	ret, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return defaultVal
	}
	return int(ret)
}

func (this *Base) MustInt64(key string, c *gin.Context) int64 {
	param := this.MustGetStrParam(key, c)
	ret, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		errInfo := fmt.Sprintf("参数[%s]非int64型，请检查", key)
		panic(ers.ParamError(errInfo))
	}
	return ret
}

func (this *Base) Int64(key string, defaultVal int64, c *gin.Context) int64 {
	param := this.getStrParam(key, c)
	if param == "" {
		return defaultVal
	}
	ret, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return defaultVal
	}
	return ret
}

func (this *Base) MustGetStrParam(key string, c *gin.Context) string {
	ret := this.getStrParam(key, c)
	if ret == "" {
		errInfo := fmt.Sprintf("参数[%s]不存在或内容为空", key)
		panic(ers.ParamError(errInfo))
	}
	return ret
}

func (this *Base) GetClientIP(c *gin.Context) string {
	//获取ip，屏蔽nginx等上层代理的情况
	return c.ClientIP()
}

func (this *Base) getStrParam(key string, c *gin.Context) string {
	// 第一步: 通过POST参数获取数据
	ret := c.PostForm(key)
	if ret == "" {
		// 第二步: 通过rest url的内部获取, 例如 /id/:id
		ret = c.Param(key)
	}
	if ret == "" {
		// 第三步: 通过url的参数获取, 例如 url?id=234
		ret = c.Query(key)
	}
	return strings.TrimSpace(ret)
}
