package conf

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

var RootPath string

var Location string

type Server struct {
	HttpPort       int
	WriteTimeout   time.Duration
	ReadTimeout    time.Duration
	RunMode        string
	MaxHeaderBytes int
	DbType         string
}

var ServerConf = &Server{
	MaxHeaderBytes: 1,
}

type File struct {
	Path               string
	MaxMultipartMemory int64
}

var FileConf = &File{
	MaxMultipartMemory: 5,
}

type Mysql struct {
	Host      string
	Port      int
	Username  string
	Password  string
	Database  string
	Charset   string
	ParseTime bool
	Location  string
}

var MysqlConf = &Mysql{
	Charset:   "utf8mb4",
	ParseTime: true,
	Location:  "Local",
}

type Pgsql struct {
	Host      string
	Port      int
	Username  string
	Password  string
	Database  string
	ParseTime string
	SSLMode   string
	TimeZone  string
}

var PgsqlConf = &Pgsql{
	SSLMode:  "disable",
	TimeZone: "Asia/Shanghai",
}

var dbServer *gorm.DB

type Redis struct {
	Host     string
	Port     int
	Password string
	Db       int
}

var (
	RedisConf   = &Redis{}
	RedisClient *redis.Client
	CtxRedis    = context.Background()
)
