package models

import (
	"Blog/src/structs"
	"database/sql"
	"fmt"
)

func ReadAccountUser(username_params string) structs.User {
	fmt.Println("ReadAccountUser model")
	var id int
	var username string
	var password string
	var createAt string
	var updateAt sql.NullString = sql.NullString{
		String: "ahihi",
		Valid:  true,
	}
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println("Error database here: ", err)
		return structs.User{}
	}

	defer db.Close()

	var statement string = "SELECT * FROM user WHERE username=$1"
	rows := db.QueryRow(statement, username_params)

	errScan := rows.Scan(&id, &username, &password, &createAt, &updateAt)

	user := structs.User{
		Id:       id,
		Username: username,
		Password: password,
		CreateAt: createAt,
		UpdateAt: updateAt,
	}

	if errScan != nil {
		fmt.Println("err Scan: ", errScan)
		// return false
		return structs.User{}
	}

	return user
}
