package db

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	VideoId      int64
	UserId       string
	VideoUrl     string
	CoverUrl     string
	Title        string
	Description  string
	VisitCount   int64
	LikeCount    int64
	CommentCount int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
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

type Counts struct {
	VisitCount   int64
	LikeCount    int64
	CommentCount int64
}
