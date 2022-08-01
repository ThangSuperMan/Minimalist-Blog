package models

import (
	"database/sql"
	"fmt"
	"log"
)

func CheckUserExists(username_params string) bool {
	var id string
	var username string
	var password string
	var createAt string
	var updateAt sql.NullString = sql.NullString{
		String: "ahihi",
		Valid:  false,
	}

	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println("Error database here: ", err)
		return false
	}

	defer db.Close()

	var statement string = "SELECT * FROM user WHERE username=$1"
	rows := db.QueryRow(statement, username_params)

	errScan := rows.Scan(&id, &username, &password, &createAt, &updateAt)

	if errScan != nil {
		fmt.Println("err Scan: ", errScan)
		return false
	}

	return true
}

func AddUser(username string, password string, createAt string) {
	fmt.Println("AddUser model")
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal("Error here: ", err)
	}

	defer db.Close()

	var query string = "INSERT INTO user (username, password, createAt) VALUES (?, ?, ?)"

	stmt, _ := db.Prepare(query)
	resultQuery, _ := stmt.Exec(username, password, createAt)
	fmt.Println("Result insert: ", resultQuery)
}
