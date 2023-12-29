package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user_system/app/constants"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Code: constants.Success,
		Msg:  constants.GetMsg(constants.Success),
		Data: data,
	}
}

func NewErrorResponse(errCode int) *Response {
	return &Response{
		Code: errCode,
		Msg:  constants.GetMsg(errCode),
	}
}

func NewErrorResponseWithError(errCode int, err error) *Response {
	return &Response{
		Code: errCode,
		Msg:  err.Error(),
	}
}

func NewErrorResponseWithData(errCode int, data ...interface{}) *Response {
	return &Response{
		Code: errCode,
		Msg:  constants.GetMsg(errCode, data...),
	}
}

func (r *Response) Return(c *gin.Context) {
	if r.Code == constants.Success {
		r.ReturnWithHttpCode(c, http.StatusOK)
	} else {
		r.ReturnWithHttpCode(c, http.StatusBadRequest)
	}
}

func (r *Response) ReturnWithHttpCode(c *gin.Context, HttpCode int) {
	c.JSON(HttpCode, &r)
}

var HttpCode int

func NewError(errCode int, error error) Response {
	var errorMsg string

	if error == nil {
		errorMsg = constants.GetMsg(errCode)
	} else {
		errorMsg = error.Error()
	}

	return Response{
		Code: errCode,
		Msg:  errorMsg,
	}
}

func NewErrorWithData(errCode int, data ...interface{}) Response {
	var errorMsg string = GetError(errCode, data...).Error()
	return Response{
		Code: errCode,
		Msg:  errorMsg,
	}
}

func NewSuccess(data interface{}) Response {
	return Response{
		Code: constants.Success,
		Msg:  constants.GetMsg(constants.Success),
		Data: data,
	}
}

func (r Response) Error(c *gin.Context) {
	if HttpCode == 0 {
		c.JSON(http.StatusBadRequest, r)
	}
	c.JSON(HttpCode, r)
}

func (r Response) Success(c *gin.Context) {
	c.JSON(http.StatusOK, r)
}

func GetError(errCode int, data ...interface{}) Error {
	return Error{Msg: constants.GetMsg(errCode, data...)}
}

type Error struct {
	Msg string
}

func (e Error) Error() string {
	return e.Msg
}
