package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"user_system/app/constants"
	"user_system/app/repository/upload"
	"user_system/app/response"
)

type FileController struct {
}

func (controller FileController) Upload(c *gin.Context) {
	r := response.Gin{C: c}
	_, headers, err := c.Request.FormFile("files")

	if err != nil {
		log.Printf("获取上传文件失败: %v", err)
		fmt.Printf("%v", err)
		r.Error(http.StatusBadRequest, constants.ErrorFileGetFail, err)
		return
	}

	filePath, errCode, err := upload.Upload(c, headers)
	if errCode != 0 {
		r.Error(http.StatusBadRequest, errCode, err)
		return
	}
	r.Success(filePath)
}
