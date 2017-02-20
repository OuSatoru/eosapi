package vals

import (
	"net/http"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Client = &http.Client{}
var Db *sql.DB
var err error

func OpenDb() {
	Db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
}