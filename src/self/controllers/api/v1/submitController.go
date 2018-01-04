package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"self/controllers/baseController"
	"self/judger"
	"self/managers"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	Result     judger.Result
}

func (this SubmitController) httpHandlerChangeSubmit(c *gin.Context) {
	buf, _ := ioutil.ReadAll(c.Request.Body)

	var mess ChangeSubMess
	if err := json.Unmarshal(buf, &mess); err != nil {
		panic(err.Error())
	}

	log.Infof("submitController: %#v", mess)

	err := managers.SubmitManager{}.ChangeSubmitResult(mess.SubmitType, mess.SubmitId, mess.Result)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.String(http.StatusOK, "OK")
}
