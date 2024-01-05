package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"user_system/utils"
)

type DB struct {
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

var MysqlConfig = &Mysql{
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

var PgsqlConfig = &Pgsql{
	SSLMode:  "disable",
	TimeZone: "Asia/Shanghai",
}

var Db *gorm.DB

func (db DB) Setup() {
	switch ServerConfig.DbType {
	case "mysql":
		utils.LoadConfig("mysql", MysqlConfig)
		fmt.Printf("%+v\n", MysqlConfig)
		mysqlConn()
	case "pgsql":
		utils.LoadConfig("pgsql", PgsqlConfig)
		pgsqlConn()
	default:
		log.Panic("数据库连接仅支持mysql,pgsql")
	}
}

func mysqlConn() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		MysqlConfig.Username,
		MysqlConfig.Password,
		MysqlConfig.Host,
		MysqlConfig.Port,
		MysqlConfig.Database,
		MysqlConfig.Charset,
		MysqlConfig.ParseTime,
		MysqlConfig.Location,
	)
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("mysql数据库连接失败:%v", err)
	}
}

func pgsqlConn() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		PgsqlConfig.Host,
		PgsqlConfig.Username,
		PgsqlConfig.Password,
		PgsqlConfig.Database,
		PgsqlConfig.Port,
		PgsqlConfig.SSLMode,
		PgsqlConfig.TimeZone,
	)
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("pgsql数据库连接失败:%v", err)
	}
}
