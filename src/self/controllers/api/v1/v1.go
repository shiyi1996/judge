/**
 * Created by shiyi on 2017/10/1.
 * Email: shiyi@fightcoder.com
 */

package controllers

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	SubmitController{}.Register(router)
}
