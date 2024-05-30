package dal

import (
	"os"
	"tiktokrpc/cmd/social/dal/db"
)

func Init() {
	err := db.Init()
	if err != nil {
		os.Exit(1)
		return
	}
}
