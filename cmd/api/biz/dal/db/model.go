package db

import "time"

type Message struct {
	MsgId      int64
	FromUserId string
	ToUserId   string
	Content    string
	CreatedAt  time.Time
	Status     int64
}
