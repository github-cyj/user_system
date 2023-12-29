package models

import (
	"gorm.io/gorm"
	"time"
	"user_system/config"
)

type BaseModel struct {
	ID         uint           `gorm:"primaryKey"  json:"id"`
	CreateTime time.Time      `gorm:"autoCreateTime" json:"create_time" go:"unix"`
	UpdateTime time.Time      `gorm:"autoUpdateTime" json:"update_time" format_time:"unix"`
	DeleteTime gorm.DeletedAt `gorm:"index;->:false" json:"-" format:"unix"`
}

func NewDb() *gorm.DB {
	return config.Db.Debug()
}
