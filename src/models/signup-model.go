package models

import (
	"database/sql"
	"fmt"
)

func AddUser(username string, password string) {
	fmt.Println("AddUser model")
}

func CheckUserExists(username_params string) bool {
	fmt.Println("username_params: ", username_params)
	fmt.Println("CheckUserExists")
	var id string
	var username string
	var password string

	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println("Error database here: ", err)
		return false
	}

	defer db.Close()

	var statement string = "SELECT * FROM user WHERE username=$1"
	rows := db.QueryRow(statement, username_params)

	errScan := rows.Scan(&id, &username, &password)

	if errScan != nil {
		fmt.Println("errorScan: ", errScan)
		return false
	}

	return true
}
