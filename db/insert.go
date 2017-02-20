package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InsUser(db *sql.DB, userName string, password string) int64 {
	stmt, err := db.Prepare("INSERT INTO login(username, password) VALUES (?, ?)")
	checkErr(err)
	res, err := stmt.Exec(userName, password)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	return id
}

func UpdJsession(db *sql.DB, userName string, jsession string) int64 {
	stmt, err := db.Prepare("UPDATE login SET jsessionid = ? WHERE username = ?")
	checkErr(err)
	res, err := stmt.Exec(jsession, userName)
	checkErr(err)
	af, err := res.RowsAffected()
	checkErr(err)
	return af
}

func UpdUnread(db *sql.DB, userName string, unread string) int64 {
	stmt, err := db.Prepare("UPDATE login SET unread = ? WHERE username = ?")
	checkErr(err)
	res, err := stmt.Exec(unread, userName)
	checkErr(err)
	af, err := res.RowsAffected()
	checkErr(err)
	return af
}

func SelJsession(db *sql.DB, userName string) string {
	row := db.QueryRow(fmt.Sprintf("SELECT jsessionid FROM login WHERE username = '%s'", userName))
	var jsession string
	err := row.Scan(&jsession)
	checkErr(err)
	return jsession
}

func SelPwd(db *sql.DB, userName string) string {
	row := db.QueryRow(fmt.Sprintf("SELECT password FROM login WHERE username = '%s'", userName))
	var pwd string
	err := row.Scan(&pwd)
	checkErr(err)
	return pwd
}

func DelUser(db *sql.DB, userName string) int64 {
	stmt, err := db.Prepare("DELETE FROM login WHERE username = ?")
	checkErr(err)
	res, err := stmt.Exec(userName)
	checkErr(err)
	af, err := res.RowsAffected()
	checkErr(err)
	return af
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
