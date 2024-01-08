package baseType

import (
	"fmt"
	"time"
)

type JsonTime time.Time

// MarshalJSON 实现它的json序列化方法
func (jsonTime JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(jsonTime).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
