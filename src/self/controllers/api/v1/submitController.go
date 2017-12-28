package controllers

import (
	"encoding/json"
	"io/ioutil"

	"self/controllers/baseController"
	"self/judger"
	"self/managers"

	"github.com/gin-gonic/gin"
)

type SubmitController struct {
	baseController.Base
}

func (this SubmitController) Register(router *gin.Engine) {
	router.POST("/change_submit", this.httpHandlerChangeSubmit)
}

type ChangeSubMess struct {
	SubmitType string
	SubmitId   int64
	result     judger.Result
}

func (this SubmitController) httpHandlerChangeSubmit(c *gin.Context) {
	buf, _ := ioutil.ReadAll(c.Request.Body)
	var mess ChangeSubMess
	if err := json.Unmarshal(buf, &mess); err != nil {
		panic(err.Error())
	}

	managers.SubmitManager{}.ChangeSubmitResult(mess.SubmitType, mess.SubmitId, mess.result)
}
