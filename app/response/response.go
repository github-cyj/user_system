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
	//if r.Code == constants.Success {
	r.ReturnWithHttpCode(c, http.StatusOK)
	//} else {
	//	r.ReturnWithHttpCode(c, http.StatusBadRequest)
	//}
}

func (r *Response) ReturnWithData(c *gin.Context, data interface{}) {
	if r == nil {
		NewSuccessResponse(data).Return(c)
	} else {
		r.Return(c)
	}
}

func (r *Response) ReturnWithHttpCode(c *gin.Context, HttpCode int) {
	c.JSON(HttpCode, &r)
}
