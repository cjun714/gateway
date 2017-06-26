package dao

import (
	"database/sql"
	"util/log"
)

func demo() {
	// sql.Open("<driver>", "<usrname>:<passwd>@tcp(<ip>:<port>)/<db>")
	db, e := sql.Open("mysql", "cjun:test@tcp(192.168.10.128:3306)/testdb")
	checkErr(e)
	defer db.Close()

	e = db.Ping() // check if connect to DB ok
	checkErr(e)

	rows, e := db.Query("select appkey,user_name,passwd from auth where appkey=?", 1)
	checkErr(e)
	defer rows.Close()
	for rows.Next() {
		usr := User{}
		e = rows.Scan(&usr.Appkey, &usr.UserName, &usr.Passwd)
		log.H(usr)
		checkErr(e)
	}

	db.Close()
}

func checkErr(e error) {
	if e != nil {
		log.E(e)
		panic(e)
	}
}

// User is used to hold user's name and passwd
type User struct {
	Appkey   int
	UserName string
	Passwd   string
}
