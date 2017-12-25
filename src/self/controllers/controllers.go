/**
 * Created by shiyi on 2017/10/1.
 * Email: shiyi@fightcoder.com
 */

package controllers

import (
	apiv1 "self/controllers/api/v1"
	authv1 "self/controllers/auth/v1"
	cself "self/controllers/self"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	apiv1.Register(router)
	authv1.Register(router)
	cself.Register(router)
}
