package config

import (
	"user_system/utils"
)

type File struct {
	Path               string
	MaxMultipartMemory int64
}

var FileConfig = &File{
	MaxMultipartMemory: 5,
}

func (file File) Setup() {
	utils.LoadConfig("file", FileConfig)
	FileConfig.MaxMultipartMemory = FileConfig.MaxMultipartMemory * 1024 * 1024
}
