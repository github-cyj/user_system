package models

import (
	"gorm.io/gorm"
	"user_system/app/baseType"
	"user_system/config"
)

type BaseModel struct {
	ID         uint              `gorm:"primaryKey"  json:"id"`
	CreateTime baseType.JsonTime `gorm:"autoCreateTime" json:"create_time" go:"unix"`
	UpdateTime baseType.JsonTime `gorm:"autoUpdateTime" json:"update_time" format_time:"unix"`
	DeleteTime gorm.DeletedAt    `gorm:"index;->:false" json:"-" format:"unix"`
}

func NewDb() *gorm.DB {
	return config.Db.Debug()
}
