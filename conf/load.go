package conf

import (
	"flag"
	"github.com/go-ini/ini"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

var conf *ini.File

func Setup() {
	RootPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	var iniPath = flag.String("c", path.Join(RootPath, "app.ini"), "ini配置文件")
	RootPath = filepath.Dir(*iniPath)
	var err error
	conf, err = ini.Load(*iniPath)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse app.ini': %v", err)
	}
	loadServeConf()
	loadFileConf()
	loadDb()
	loadRedis()
}

func loadSection(section string, v interface{}) {
	err := conf.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("conf load %s err: %v", section, err)
	}

}

func loadServeConf() {
	loadSection("server", ServerConf)
	ServerConf.ReadTimeout = ServerConf.ReadTimeout * time.Second
	ServerConf.WriteTimeout = ServerConf.WriteTimeout * time.Second
	ServerConf.MaxHeaderBytes = ServerConf.MaxHeaderBytes * 1024 * 1024
}

func loadFileConf() {
	loadSection("file", FileConf)
	FileConf.MaxMultipartMemory = FileConf.MaxMultipartMemory * 1024 * 1024
}

func loadDb() {
	switch ServerConf.DbType {
	case "mysql":
		loadSection("mysql", MysqlConf)
	case "pgsql":
		loadSection("pgsql", PgsqlConf)
	default:
		log.Panic("暂不支持此数据库，请联系管理员")
	}
}

func loadRedis() {
	loadSection("redis", RedisConf)
}
