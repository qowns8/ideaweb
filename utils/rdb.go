package utils

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"os"
)

type RDB struct {
	Db * sql.DB
	user string
	root string
	pwd string
	db_name string
}


func NewRDB() *RDB{

	var rdb RDB
	root := os.Getenv("GOAPP_MYSQL")

	db, _ := sql.Open("mysql", root) //"root:pwd@tcp(127.0.0.1:3306)/testdb")

	rdb.Db = db
	return &rdb
}
