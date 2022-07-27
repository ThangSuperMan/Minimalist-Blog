package models

import (
	"Blog/src/structs"
	"database/sql"
	"fmt"
	"strconv"
)

func QueryABlog(id string) structs.Blog {
	fmt.Println("QueryABlog")
	var title string
	var content string

	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println("error here: ", err)
		return structs.Blog{}
	}

	defer db.Close()

	intVar, err := strconv.Atoi(id)

	var statement string = `SELECT title, content FROM blogs WHERE id=?`
	row := db.QueryRow(statement, intVar)

	errScan := row.Scan(&title, &content)

	if errScan != nil {
		fmt.Println("Got an error when trying to select data")
		fmt.Println("Error here: ", errScan)
		return structs.Blog{}
	}

	blog := structs.Blog{
		Title: title, Content: content,
	}

	return blog
}
