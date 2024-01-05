package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"user_system/base"
)

var cfg *ini.File

func LoadConfig(section string, v interface{}) {
	var err error
	cfg, err = ini.Load(base.IniPath)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse app.ini': %v", err)
	}
	err = cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}
