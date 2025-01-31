package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int32     `gorm:"type:int;primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(255);not null"`
	NickName string     `gorm:"type:varchar(25);"`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female=女,male=男'"`
	Role     int        `gorm:"column:role;default:1;type:int comment '1=普通用户 2=管理员'"`
}
