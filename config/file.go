package config

import (
	"user_system/src/utils"
)

type File struct {
	Path               string
	MaxMultipartMemory int64
}

var FileConfig = &File{
	MaxMultipartMemory: 5 * 1024 * 1024,
}

func (file File) Setup() {
	utils.LoadConfig("file", FileConfig)
}
