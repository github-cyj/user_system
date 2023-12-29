package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"user_system/config"
	"user_system/routes"
)

func init() {
	config.Server{}.Setup()
	config.DB{}.Setup()
	config.File{}.Setup()
}

func main() {
	// 日志输出到文件和控制台两个位置
	// gin.DefaultWriter = io.MultiWriter(RefreshLogFileUsage(), os.Stdout)

	gin.SetMode(config.ServerConfig.RunMode)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.ServerConfig.HttpPort),
		Handler:        routes.Route(),
		ReadTimeout:    config.ServerConfig.ReadTimeout,
		WriteTimeout:   config.ServerConfig.WriteTimeout,
		MaxHeaderBytes: config.ServerConfig.MaxHeaderBytes,
	}
	log.Printf("[info] start http server listening :%d", config.ServerConfig.HttpPort)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}

//func RefreshLogFileUsage() *os.File {
//	var targetFile *os.File
//	fileName := time.Now().Format("2006-04-02_15_04") + ".log"
//
//	tryFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0777)
//	if err != nil {
//		_ = os.WriteFile(fileName, []byte(""), 0777)
//		targetFile, _ = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
//	} else {
//		targetFile = tryFile
//	}
//	return targetFile
//}
