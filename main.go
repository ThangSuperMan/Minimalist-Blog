package main

import (
	"Blog/src/models"
	"Blog/src/routers"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Blog struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p>homepage</p>")
}

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	router := mux.NewRouter()

	if err != nil {
		log.Fatal("Error here", err)
		fmt.Println("hello")
	}

	defer db.Close()

	var statement string = `
	CREATE TABLE IF NOT EXISTs "blogs" (
	"id"	INTEGER NOT NULL,
	"title"	TEXT NOT NULL,
	"content"	INTEGER NOT NULL,
	PRIMARY KEY("id")
	);`

	result, err := db.Exec(statement)
	if err != nil {
		log.Fatal("\n Error here: ", err)
	}

	fmt.Println("Result: ", result)

	dir := http.Dir("./static")
	fs := http.FileServer(dir)
	router.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", fs))

	router.HandleFunc("/", routers.HandleHomeRouter).Methods("GET")
	router.HandleFunc("/detail_blog/{id}", routers.HandleDetaiBloglRouter).Methods("GET")
	router.HandleFunc("/add_blog", routers.HandleAddBlogRouter).Methods("GET")
	router.HandleFunc("/add_blog", models.HandleAddBlog).Methods("POST")
	router.HandleFunc("/login", routers.HandleLoginRouter).Methods("GET")

	http.ListenAndServe(":3000", router)
}
