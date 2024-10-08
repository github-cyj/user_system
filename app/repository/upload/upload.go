package upload

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"user_system/app/constants"
	"user_system/app/response"
	"user_system/conf"
)

func Upload(c *gin.Context, headers *multipart.FileHeader) (saveFilePath string, r *response.Response) {
	if headers.Size > conf.FileConf.MaxMultipartMemory {
		maxSize := conf.FileConf.MaxMultipartMemory / 1024 / 1024
		log.Printf("文件大于%dM", maxSize)
		return saveFilePath, response.NewErrorResponseWithData(constants.ErrorFileExceedsMaxSize, maxSize)
	}

	md5Str, r := GetFileMD5ByHeaders(headers)
	if r != nil {
		return saveFilePath, r
	}
	fileExt := filepath.Ext(headers.Filename)

	saveFilePath = conf.FileConf.Path + md5Str + fileExt
	_, err := os.Stat(saveFilePath)
	// 文件不存在
	if os.IsNotExist(err) {
		err = c.SaveUploadedFile(headers, saveFilePath)
		if err != nil {
			log.Printf("文件上传失败")
			return "", response.NewErrorResponse(constants.ErrorFileUploadFail)
		}
	}
	return saveFilePath, r
}

func GetFileMD5ByPath(pathName string) (md5Str string, r *response.Response) {
	f, err := os.Open(pathName)
	if err != nil {
		log.Printf("文件打开失败:%s", err)
		r = response.NewErrorResponse(constants.ErrorFileOpenFail)
		return md5Str, r
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	md5hash := md5.New()
	if _, err = io.Copy(md5hash, f); err != nil {
		log.Printf("复制副本失败:%s", err)
		r = response.NewErrorResponse(constants.ErrorFileCopyFail)
		return md5Str, r
	}
	has := md5hash.Sum(nil)
	md5Str = fmt.Sprintf("%x", has)
	return md5Str, r
}

func GetFileMD5ByHeaders(headers *multipart.FileHeader) (md5Str string, r *response.Response) {
	f, err := headers.Open()
	if err != nil {
		log.Printf("文件打开失败:%s", err)
		r = response.NewErrorResponse(constants.ErrorFileOpenFail)
		return md5Str, r
	}
	defer func(f multipart.File) {
		_ = f.Close()
	}(f)

	md5hash := md5.New()
	if _, err = io.Copy(md5hash, f); err != nil {
		fmt.Println("Copy", err)
		r = response.NewErrorResponse(constants.ErrorFileCopyFail)
		return md5Str, r
	}
	has := md5hash.Sum(nil)
	md5Str = fmt.Sprintf("%x", has)
	return md5Str, r
}
