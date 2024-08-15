package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"user_system/app/baseType"
	"user_system/conf"
)

type BaseModel struct {
	ID         uint              `gorm:"primaryKey"  json:"id"`
	CreateTime baseType.JsonTime `gorm:"autoCreateTime" json:"create_time" go:"unix"`
	UpdateTime baseType.JsonTime `gorm:"autoUpdateTime" json:"update_time" format_time:"unix"`
	DeleteTime gorm.DeletedAt    `gorm:"index;->:false" json:"-" format:"unix"`
}

var Db *gorm.DB

func Setup() {
	switch conf.ServerConf.DbType {
	case "mysql":
		Db = mysqlConn().Debug()
	case "pgsql":
		Db = pgsqlConn().Debug()
	default:
		log.Panicf("暂不支持此数据库，请联系管理员")
	}
}

func mysqlConn() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		conf.MysqlConf.Username,
		conf.MysqlConf.Password,
		conf.MysqlConf.Host,
		conf.MysqlConf.Port,
		conf.MysqlConf.Database,
		conf.MysqlConf.Charset,
		conf.MysqlConf.ParseTime,
		conf.MysqlConf.Location,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("mysql数据库连接失败:%v", err)
	}
	return db
}

func pgsqlConn() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		conf.PgsqlConf.Host,
		conf.PgsqlConf.Username,
		conf.PgsqlConf.Password,
		conf.PgsqlConf.Database,
		conf.PgsqlConf.Port,
		conf.PgsqlConf.SSLMode,
		conf.PgsqlConf.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("pgsql数据库连接失败:%v", err)
	}
	return db

}
