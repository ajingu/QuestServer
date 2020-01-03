package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)
var db *sql.DB

func init(){
	var err error
	db, err = sql.Open("postgres", "user=gwp dbname=quest password=gwp sslmode=disable")
	if err != nil{
		panic(err)
	}
}

func retrieveAll()(users []User, err error){
	rows, err := db.Query("SELECT id, name, is_paid FROM users")
	if err != nil{
		return
	}

	for rows.Next(){
		user := User{}
		err = rows.Scan(&user.Id, &user.Name, &user.IsPaid)
		if err != nil{
			return
		}

		users = append(users, user)
	}
	rows.Close()

	return
}
