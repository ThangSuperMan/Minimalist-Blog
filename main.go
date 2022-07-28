package main

import (
	"Blog/src/controllers"
	"Blog/src/models"
	"Blog/src/routers"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func initDatabase() {
	fmt.Println("initDatabase")
	// fmt.Println("Current date: ", time.Now())
	currentTime := time.Now()

	fmt.Println(currentTime)

	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal("Error here: ", err)
	}

	defer db.Close()

	// var statement string = `drop table user;`

	var statement string = `
		CREATE TABLE IF NOT EXISTS  "user" (
		"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username"	TEXT NOT NULL,
		"password"	TEXT NOT NULL,
		"createAt"  TEXT NOT NULL,
		"updateAt"	TEXT
	);

		CREATE TABLE IF NOT EXISTS "blogs" (
		"id"	INTEGER NOT NULL,
		"title"	TEXT NOT NULL,
		"content"	INTEGER NOT NULL,
		PRIMARY KEY("id")
	);
	`

	db.Exec(statement)
}

func main() {
	initDatabase()
	router := mux.NewRouter()
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
