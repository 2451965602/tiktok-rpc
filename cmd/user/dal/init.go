package dal

import (
	"os"
	"tiktokrpc/cmd/user/dal/db"
)

func Init() {
	err := db.Init()
	if err != nil {
		os.Exit(1)
		return
	}
}
