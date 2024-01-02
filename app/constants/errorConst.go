package constants

import (
	"fmt"
	"strings"
)

const (
	Success = 200
	Error   = 500

	ErrorBind               = 10404001
	ErrorFileExceedsMaxSize = 10500001
	ErrorFileUploadFail     = 10500002
	ErrorFileOpenFail       = 10500003
	ErrorFileCopyFail       = 10500004
	ErrorFileGetFail        = 10500005
	ErrorNotExits           = 10500006
	ErrorAddFail            = 10500007
)

var MsgMap = map[int]string{
	Success:                 "success",
	Error:                   "error",
	ErrorBind:               "参数获取失败",
	ErrorFileExceedsMaxSize: "文件大于%vM",
	ErrorFileUploadFail:     "文件上传失败",
	ErrorFileOpenFail:       "文件打开失败",
	ErrorFileCopyFail:       "文件复制副本失败",
	ErrorFileGetFail:        "获取上传文件失败: %v",
	ErrorNotExits:           "%s不存在",
	ErrorAddFail:            "%s添加失败",
}

func GetMsg(errorCode int, data ...interface{}) string {
	msg, ok := MsgMap[errorCode]
	if ok {
		strCount := strings.Count(msg, "%")
		if strCount == 0 {
			return msg
		} else {
			dataCount := len(data)

			if dataCount == strCount {
				return fmt.Sprintf(msg, data...)
			} else {
				return fmt.Sprint(msg, data, "通过%检测，参数数量与提供内容不符")
			}
		}
	}
	return MsgMap[Error]
}
