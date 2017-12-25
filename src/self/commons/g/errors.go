package g

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// fatal
type Fatal struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func FatalError(msg string, data ...interface{}) Fatal {
	ret := Fatal{Code: http.StatusInternalServerError, Msg: msg}
	if len(data) != 0 {
		ret.Data = data[0]
	}
	return ret
}

// error
type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func DBError(msg ...string) Error {
	return _build(http.StatusInternalServerError, "DB Error", msg...)
}

func ParamError(msg ...string) Error {
	return _build(http.StatusBadRequest, "Param Error", msg...)
}

func ServerError(msg ...string) Error {
	return _build(http.StatusInternalServerError, "Server Error", msg...)
}

func PrivError(msg ...string) Error {
	return _build(http.StatusForbidden, "Forbidden", msg...)
}

//
func NotFoundError(msg ...string) Error {
	return _build(http.StatusNotFound, "Not Found", msg...)
}

func RequestError(msg ...string) Error {
	return _build(http.StatusBadRequest, "Bad Request", msg...)
}

func HandleError() {
	if err := recover(); err != nil {
		if msg, ok := err.(Error); ok {
			log.Print(msg)
			//我们自己程序内部触发的panic
		} else {
			if err_2, ok := err.(error); ok {
				//request的异常的panic
				msg := err_2.Error()
				log.Print(msg)
			} else {
			}
		}
	}
}

func _build(code int, defval string, custom ...string) Error {
	msg := defval
	if len(custom) > 0 {
		msg = custom[0]
	}

	return Error{
		Code: code,
		Msg:  msg,
	}
}
