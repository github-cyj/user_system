package baseType

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type JsonTime time.Time

// MarshalJSON 实现它的json序列化方法
func (jsonTime JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(jsonTime).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (jsonTime *JsonTime) Scan(value interface{}) error {
	*jsonTime = JsonTime(value.(time.Time))
	return nil
}

func (jsonTime JsonTime) Value() (driver.Value, error) {
	return time.Time(jsonTime), nil
}
