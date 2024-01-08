package models

type User struct {
	BaseModel
	Username string `gorm:"size:32" json:"username"`
	Avatar   string `form:"avatar" json:"avatar"`
}

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "user"
}
