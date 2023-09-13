package utils

import (
	"github.com/go-ini/ini"
	"log"
)

var cfg *ini.File

func LoadConfig(section string, v interface{}) {
	var err error
	cfg, err = ini.Load(".ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	err = cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
