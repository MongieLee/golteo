package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDb() (err error) {
	dsn := "root:bendimima@tcp(127.0.0.1:3306)/go-pure"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = Db.Ping()
	if err != nil {
		return err
	}
	return nil
}
