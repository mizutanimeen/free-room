package db

import (
	"database/sql"
	"fmt"
	"os"
)

func Open() (*sql.DB, error) {
	tDbPath := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_IP"), os.Getenv("MYSQL_DATABASE"))
	tDb, tError := sql.Open("mysql", tDbPath)
	if tError != nil {
		return nil, tError
	}
	if tError := tDb.Ping(); tError != nil {
		return nil, fmt.Errorf("データベース接続失敗: %s\n", tError.Error())
	}

	return tDb, nil
}
