package config

import (
	"time"
	"user_system/utils"
)

type Server struct {
	HttpPort       int
	WriteTimeout   time.Duration
	ReadTimeout    time.Duration
	RunMode        string
	MaxHeaderBytes int
	DbType         string
}

var ServerConfig = &Server{
	MaxHeaderBytes: 1,
}

func (server Server) Setup() {
	utils.LoadConfig("server", ServerConfig)
	ServerConfig.ReadTimeout = ServerConfig.ReadTimeout * time.Second
	ServerConfig.WriteTimeout = ServerConfig.WriteTimeout * time.Second
	ServerConfig.MaxHeaderBytes = ServerConfig.MaxHeaderBytes * 1024 * 1024
}
