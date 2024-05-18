package dal

import "tiktokrpc/cmd/user/dal/db"

func init() {
	err := db.Init()
	if err != nil {
		return
	}
}
