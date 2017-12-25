/**
 * Created by shiyi on 2017/10/1.
 * Email: shiyi@fightcoder.com
 */

package controllers

import (
	"net/http"
	"self/commons/g"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	router.GET("/self/health", checkHealth)
	router.GET("/self/config", showConfig)
	router.GET("/self/version", showVersion)
}

//返回健康检查
func checkHealth(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

//显示配置
func showConfig(c *gin.Context) {
	c.JSON(http.StatusOK, g.Conf())
}

//显示版本
func showVersion(c *gin.Context) {
	c.String(http.StatusOK, "ver: %s\nbuildTime: %s\n", g.Version(), g.BuildInfo())
}
