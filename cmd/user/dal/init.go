package dal

import (
	"os"
	"tiktokrpc/cmd/user/dal/db"
	milvus "tiktokrpc/cmd/user/dal/miluvs"
)

func Init() {
	err := db.Init()
	if err != nil {
		os.Exit(1)
		return
	}

	milvus.Init()

}
