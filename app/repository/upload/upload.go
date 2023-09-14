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
	"user_system/config"
)

func Upload(c *gin.Context, headers *multipart.FileHeader) (filePath string, errCode int, err error) {
	if headers.Size > config.FileConfig.MaxMultipartMemory {
		maxSize := config.FileConfig.MaxMultipartMemory >> 20
		log.Printf("文件大于%dM", maxSize)
		return "", constants.ErrorFileExceedsMaxSize, response.Error{fmt.Sprintf("%d", maxSize)}
	}

	md5Str, errCode, err := GetFileMD5ByHeaders(headers)
	if errCode != 0 {
		return "", errCode, err
	}
	fileExt := filepath.Ext(headers.Filename)
	saveFilePath := config.FileConfig.Path + md5Str + fileExt
	_, err = os.Stat(saveFilePath)
	// 文件不存在
	if os.IsNotExist(err) {
		err = c.SaveUploadedFile(headers, saveFilePath)
		if err != nil {
			log.Printf("文件上传失败")
			return "", constants.ErrorFileUploadFail, err
		}
	}
	return saveFilePath, 0, nil
}

func GetFileMD5ByPath(pathName string) (md5Str string, errCode int, err error) {
	f, err := os.Open(pathName)
	if err != nil {
		log.Printf("文件打开失败:%s", err)
		return "", constants.ErrorFileOpenFail, err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	md5hash := md5.New()
	if _, err = io.Copy(md5hash, f); err != nil {
		log.Printf("复制副本失败:%s", err)
		return "", constants.ErrorFileCopyFail, err

	}
	has := md5hash.Sum(nil)
	md5Str = fmt.Sprintf("%x", has)
	return md5Str, 0, err
}

func GetFileMD5ByHeaders(headers *multipart.FileHeader) (md5Str string, errCode int, err error) {
	f, err := headers.Open()
	if err != nil {
		log.Printf("文件打开失败:%s", err)
		return "", constants.ErrorFileOpenFail, err
	}
	defer func(f multipart.File) {
		_ = f.Close()
	}(f)

	md5hash := md5.New()
	if _, err = io.Copy(md5hash, f); err != nil {
		fmt.Println("Copy", err)
		return "", constants.ErrorFileCopyFail, err
	}
	has := md5hash.Sum(nil)
	md5Str = fmt.Sprintf("%x", has)
	return md5Str, 0, nil
}
