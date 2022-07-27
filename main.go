package main

import (
	"Blog/src/controllers"
	"Blog/src/models"
	"Blog/src/routers"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	router := mux.NewRouter()

	if err != nil {
		log.Fatal("Error here: ", err)
	}

	defer db.Close()

	// var statement string = `drop table user`

	// var statement string = `
	// CREATE TABLE IF NOT EXISTS "blogs" (
	// "id"	INTEGER NOT NULL,
	// "title"	TEXT NOT NULL,
	// "content"	INTEGER NOT NULL,
	// PRIMARY KEY("id")
	// );`

	var statement string = `
		CREATE TABLE IF NOT EXISTS  "user" (
		"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username"	TEXT NOT NULL,
		"password"	TEXT NOT NULL
	);`

	result, err := db.Exec(statement)
	if err != nil {
		// log.Fatal("Error her: ", err)
		log.Fatal("Error here: ", err)
	}

	fmt.Println("Result: ", result)

	dir := http.Dir("./static")
	fs := http.FileServer(dir)
	router.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", fs))

	router.HandleFunc("/", routers.HandleHomeRouter).Methods("GET")
	router.HandleFunc("/login", routers.HandleLoginRouter).Methods("GET")
	router.HandleFunc("/detail_blog/{id}", routers.HandleDetaiBloglRouter).Methods("GET")
	router.HandleFunc("/add_blog", routers.HandleAddBlogRouter).Methods("GET")
	router.HandleFunc("/add_blog", models.HandleAddBlog).Methods("POST")
	router.HandleFunc("/signup", routers.HandleSignUpRouter).Methods("GET")
	router.HandleFunc("/signup", controllers.SignupAccount).Methods("POST")

	http.ListenAndServe(":3000", router)
}
