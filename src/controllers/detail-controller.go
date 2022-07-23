package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type Blog struct {
	Id      int
	Title   string
	Content string
}

func RenderDetailPage(w http.ResponseWriter, r *http.Request) Blog {
	fmt.Println("controller RenderDetailPage")
	vars := mux.Vars(r)
	// var id string = vars["id"]
	// var id int = vars["id"]
	id := vars["id"]
	var title string
	var content string

	fmt.Println("Vars here: ", vars)
	fmt.Println("Id here: ", id)
	fmt.Println("Id type is: ", reflect.TypeOf(id))

	t, err := template.ParseGlob("./src/views/*.html")

	if err != nil {
		fmt.Println("Error here", err)
		return Blog{}
	}

	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println("Error database here: ", err)
		return Blog{}
	}

	defer db.Close()

	fmt.Println(id)
	var statement string = `SELECT title, content FROM blogs WHERE id=?`
	// row := db.QueryRow(`SELECT * FROM blogs WHERE id = ?`, id)
	row := db.QueryRow(statement, id)

	fmt.Println("My row: ", row)
	errScan := row.Scan(&title, &content)
	fmt.Println("Error scan: ", errScan)

	if err != nil {
		fmt.Println("Got an error when trying to select data")
		fmt.Println("Error here: ", err)
		return Blog{}
	}

	fmt.Println("Id here: ", id)
	fmt.Println("Title here: ", title)
	fmt.Println("Content here: ", content)

	intVar, err := strconv.Atoi(id)

	blog := Blog{
		Id:      intVar,
		Title:   title,
		Content: content,
	}

	t.ExecuteTemplate(w, "detail_blog.html", blog)

	return Blog{}
}
