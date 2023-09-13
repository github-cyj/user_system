package config

import (
	"user_system/pkg/utils"
)

type File struct {
	Path               string
	MaxMultipartMemory int64
}

var FileConfig = &File{
	MaxMultipartMemory: 8 << 20,
}

func init() {
	utils.LoadConfig("file", FileConfig)
}
