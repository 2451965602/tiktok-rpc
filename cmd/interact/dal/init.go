package dal

import (
	"os"
	"tiktokrpc/cmd/interact/dal/db"
	"tiktokrpc/cmd/interact/dal/redis"
)

func Init() {
	err := db.Init()
	if err != nil {
		os.Exit(1)
		return
	}

	err = redis.Init()
	if err != nil {
		os.Exit(2)
		return
	}
}
