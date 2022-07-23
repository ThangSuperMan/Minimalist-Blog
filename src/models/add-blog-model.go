package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func HandleAddBlog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	blog := Blog{}
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	blog.Title = r.Form.Get("title")
	blog.Content = r.Form.Get("content")
	fmt.Printf("title %v, content %v", blog.Title, blog.Content)

	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal("\nError here: ", err)
	}
	var query string = "INSERT INTO blogs (title, content) VALUES (?, ?)"

	stmt, _ := db.Prepare(query)
	result, _ := stmt.Exec(blog.Title, blog.Content)
	fmt.Println("Result", result)

	defer db.Close()

	http.Redirect(w, r, "/", http.StatusFound)
}
