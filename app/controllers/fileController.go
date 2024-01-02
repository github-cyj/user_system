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

func (controller FileController) UploadSingle(c *gin.Context) {
	_, headers, err := c.Request.FormFile("files")
	if err != nil {
		log.Printf("获取上传文件失败: %v", err)
		response.NewErrorResponseWithError(constants.ErrorFileGetFail, err).Return(c)
		return
	}
	saveFilePath, r := upload.Upload(c, headers)
	r.ReturnWithData(c, saveFilePath)
}

func (controller FileController) UploadMulti(c *gin.Context) {
	form, err := c.MultipartForm()

	if err != nil {
		log.Printf("获取上传文件失败: %v", err)
		response.NewErrorResponseWithError(constants.ErrorFileGetFail, err).Return(c)
		return
	}
	files := form.File["file"]
	if len(files) == 0 {
		log.Printf("获取上传文件失败: %v", "file不存在")
		response.NewErrorResponseWithData(constants.ErrorNotExits, "file").Return(c)
		return
	}
	var filePathList []string
	for index, headers := range files {
		filePath, r := upload.Upload(c, headers)
		filePathList = append(filePathList, filePath)
		if r != nil {
			log.Printf("第%d张图片上传错误,错误信息为：%v", index, r.Msg)
			r.Return(c)
			return
		}
	}
	response.NewSuccessResponse(filePathList).Return(c)
}
