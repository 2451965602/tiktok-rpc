package db

import (
	"gorm.io/gorm"
	"time"
)

type Like struct {
	UserId  int64
	RootId  int64
	VideoId int64
}

type Comment struct {
	CommentId int64
	UserId    string
	VideoId   string
	RootId    string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

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
