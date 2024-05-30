package dal

import (
	"os"
	"tiktokrpc/cmd/video/dal/db"
	"tiktokrpc/cmd/video/dal/redis"
)

func Init() {
	err := db.Init()
	if err != nil {
		os.Exit(1)
		return
	}

	err = redis.Init()
	if err != nil {
		os.Exit(1)
		return
	}
}
