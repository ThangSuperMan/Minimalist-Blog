package main

import (
	"Blog/src/routers"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

type Blog struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p>homepage</p>")
}

func handleAddBlog(w http.ResponseWriter, r *http.Request) {
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

	// Save data
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
func addBlogPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("addBlogPage get")
	t, err := template.ParseFiles("./src/views/add_blog.html")
	if err != nil {
		// log.Fatal("Error here: ")
		fmt.Println("Error here: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("hello")
		return
	}

	t.Execute(w, nil)
}

func main() {
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal("Error here", err)
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

	// Serve static file

	router := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	router.HandleFunc("/", routers.HandleHomeRouter).Methods("GET")
	router.HandleFunc("/add_blog", addBlogPage).Methods("GET")
	router.HandleFunc("/add_blog", handleAddBlog).Methods("POST")
	http.ListenAndServe(":3000", router)
}
