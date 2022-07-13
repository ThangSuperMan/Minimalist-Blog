package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Blog struct {
	Id      int
	Title   string
	Content string
}

func ReadAllBlogs() []Blog {
	fmt.Println("ReadAllBlogs")

	var id int
	var title string
	var content string

	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println("Error database here: ", err)
		return nil
	}

	defer db.Close()

	var statement string = "SELECT * FROM blogs"

	rows, err := db.Query(statement)

	blogs := make([]Blog, 0)

	for rows.Next() {
		rows.Scan(&id, &title, &content)

		blog := Blog{
			Id:      id,
			Title:   title,
			Content: content,
		}

		blogs = append(blogs, blog)
	}

	return blogs
}
