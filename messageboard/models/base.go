package models

import (
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(host.docker.internal:3306)/testdb")
	if err != nil {
		fmt.Println(err)
		return
	}

	Db = database
}