package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"user_system/app/cache"
	"user_system/app/models"
	"user_system/conf"
	"user_system/routes"
)

func init() {
	// 配置加载
	conf.Setup()
	models.Setup()
	cache.Setup()

}

func main() {
	// 日志输出到文件和控制台两个位置
	gin.DefaultWriter = io.MultiWriter(RefreshLogFileUsage(), os.Stdout)

	gin.SetMode(conf.ServerConf.RunMode)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.ServerConf.HttpPort),
		Handler:        routes.Route(),
		ReadTimeout:    conf.ServerConf.ReadTimeout,
		WriteTimeout:   conf.ServerConf.WriteTimeout,
		MaxHeaderBytes: conf.ServerConf.MaxHeaderBytes,
	}
	log.Printf("[info] start http server listening :%d", conf.ServerConf.HttpPort)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}

func RefreshLogFileUsage() *os.File {
	var targetFile *os.File
	var logPath = filepath.Join(conf.RootPath, "logs")
	fmt.Println("日志输出目录:" + logPath)
	_, err := os.Stat(logPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(logPath, 0777)
		if err != nil {
			fmt.Println("日志文件目录创建失败")
		}
	}
	fileName := time.Now().Format("2006-01-02") + ".log"
	fileName = filepath.Join(conf.RootPath, "logs", fileName)
	tryFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		_ = os.WriteFile(fileName, []byte(""), 0777)
		targetFile, _ = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	} else {
		targetFile = tryFile
	}
	return targetFile
}
