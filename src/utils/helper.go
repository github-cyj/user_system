package utils

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

var cfg *ini.File

func LoadConfig(section string, v interface{}) {
	var err error
	dir, _ := os.Getwd()
	cfg, err = ini.Load(dir + "/.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	err = cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
