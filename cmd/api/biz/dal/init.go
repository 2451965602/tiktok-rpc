package dal

import "tiktokrpc/cmd/api/biz/dal/db"

func MysqlInit() error {
	err := db.Init()
	if err != nil {
		return err
	}

	return nil
}
