package base

import (
	"flag"
	"os"
	"path/filepath"
)

var RootDir string

var IniPath string

func Setup() {
	RootDir, _ = os.Getwd()
	IniPath = filepath.Join(RootDir, "app.ini")
	flag.StringVar(&IniPath, "ini-path", IniPath, "ini配置文件")
	flag.Parse()
}
