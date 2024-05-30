package db

import (
	"gorm.io/gorm"
	"time"
)

type Social struct {
	UserId   int64
	ToUserId int64
	Status   int64
}

type UserInfo struct {
	UserId    int64
	Username  string
	AvatarUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserInfoDetail struct {
	UserId    int64
	Username  string
	AvatarUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	UserId    int64
	Username  string
	Password  string
	AvatarUrl string
	OptSecret string
	MfaStatus string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
