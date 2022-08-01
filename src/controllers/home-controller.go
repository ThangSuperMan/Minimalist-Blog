package controllers

import (
	"Blog/src/models"
	"Blog/src/structs"
	"fmt"
	"html/template"
	"net/http"
)

type shipData struct {
	Blogs    []structs.Blog
	Username string
	IsUser   bool
}

func RenderHomePage(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./src/views/*.html")

	if err != nil {
		fmt.Println("Error here", err)
		return
	}

	if models.ReadAllBlogs() != nil {
		ship := shipData{
			Blogs:    models.ReadAllBlogs(),
			Username: "guest",
			IsUser:   false,
		}

		templ.ExecuteTemplate(w, "index.html", ship)
	} else {
		fmt.Println("Something wrong!")
		w.WriteHeader(http.StatusInternalServerError)
		templ.ExecuteTemplate(w, "index.html", nil)
	}

}
