package constants

import "fmt"

const (
	Success = 200
	Error   = 500

	ErrorBind               = 10404001
	ErrorFileExceedsMaxSize = 10500001
	ErrorFileUploadFail     = 10500002
	ErrorFileOpenFail       = 10500003
	ErrorFileCopyFail       = 10500004
	ErrorFileGetFail        = 10500005
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
}

func GetMsg(errorCode int, data ...interface{}) string {
	msg, ok := MsgMap[errorCode]
	if ok {
		return msg
	}
	fmt.Printf("%v", data)
	return fmt.Sprintf(MsgMap[Error], data)
}
