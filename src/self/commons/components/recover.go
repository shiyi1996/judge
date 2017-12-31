package components

import (
	"net/http"

	"self/commons/g"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Recovery() gin.HandlerFunc {
	return RecoveryWithWriter()
}

func RecoveryWithWriter() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovery -> %s\n%s\n", err)
				// TODO
				// 发送邮件至Admin
				// 返回500及具体的错误详情
				if msg, ok := err.(g.Error); ok {
					//我们自己程序内部触发的panic
					http.Error(c.Writer, msg.Msg, msg.Code)
					log.Print(msg)
				} else if msg, ok := err.(g.Fatal); ok {
					http.Error(c.Writer, msg.Msg, msg.Code)
					log.Print(msg)
					// TODO 严重错误, 需要通知管理员
				} else {
					http.Error(c.Writer, "服务器内部错误", http.StatusInternalServerError)
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}
