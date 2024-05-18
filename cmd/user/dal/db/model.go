package db

import (
	"gorm.io/gorm"
	"time"
)

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

type MFA struct {
	Secret string
	Qrcode string
}
