package baseType

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type JsonTime time.Time

// MarshalJSON 实现它的json序列化方法
func (t JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

type JsonDeleteTime gorm.DeletedAt

// MarshalJSON 实现它的json序列化方法
func (t JsonDeleteTime) MarshalJSON() ([]byte, error) {
	if t.Valid {
		var stamp = fmt.Sprintf("\"%s\"", t.Time.Format("2006-01-02 15:04:05"))
		return []byte(stamp), nil
	}
	return json.Marshal(nil)
}
