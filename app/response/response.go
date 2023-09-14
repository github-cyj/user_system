package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user_system/app/constants"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, code int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: code,
		Msg:  constants.GetMsg(code),
		Data: data,
	})
}

func (g *Gin) Success(data interface{}) {
	g.Response(http.StatusOK, constants.Success, data)
}

func (g *Gin) Error(errCode int, data ...interface{}) {
	g.C.JSON(http.StatusBadRequest, Response{
		Code: errCode,
		Msg:  constants.GetMsg(errCode, data...),
	})
}

type Error struct {
	Msg string
}

func (e Error) Error() string {
	return e.Msg
}
