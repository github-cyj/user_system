package config

import (
	"time"
	"user_system/pkg/utils"
)

type Server struct {
	HttpPort       int
	WriteTimeout   time.Duration
	ReadTimeout    time.Duration
	RunMode        string
	MaxHeaderBytes int
}

var ServerConfig = &Server{
	MaxHeaderBytes: 1 << 20,
}

func init() {
	utils.LoadConfig("server", ServerConfig)
	ServerConfig.ReadTimeout = ServerConfig.ReadTimeout * time.Second
	ServerConfig.WriteTimeout = ServerConfig.WriteTimeout * time.Second
}
