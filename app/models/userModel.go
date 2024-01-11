package models

type User struct {
	BaseModel
	Username string `gorm:"size:32" json:"username"`
	Avatar   string `form:"avatar" json:"avatar"`
	Sex      uint16 `form:"sex" json:"sex"`
	Tel      string `form:"tel" json:"tel"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"-"`
}

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "user"
}
