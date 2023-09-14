package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
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
		r.Error(constants.ErrorFileGetFail, err)
		return
	}

	filePath, errCode, err := upload.Upload(c, headers)
	if errCode != 0 {
		r.Error(errCode, err)
		return
	}
	r.Success(filePath)
}
