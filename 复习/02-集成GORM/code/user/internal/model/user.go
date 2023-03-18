package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id            int
	Account       string
	Password      string
	UserName      string `gorm:"column:username"`
	EMail         string `gorm:"column:email"`
	Mobile        string
	RoleId        int
	LastLoginTime time.Time
	Status        string
	Sort          int
	CreatedTime   time.Time `gorm:"autoCreateTime"`
	UpdatedTime   time.Time `gorm:"autoUpdateTime"`
}

func (u *User) FindByName(db *gorm.DB) error {
	return db.Where(u).First(u).Error
}
