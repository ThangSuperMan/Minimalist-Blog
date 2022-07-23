package controllers

import (
	"Blog/src/models"
	"fmt"
	"html/template"
	"net/http"
)

type shipData struct {
	Blogs []models.Blog
}

func RenderHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("render homepage")

	templ, err := template.ParseGlob("./src/views/*.html")

	if err != nil {
		fmt.Println("Error here", err)
		return
	}

	if models.ReadAllBlogs() != nil {
		fmt.Println("blogs: ", models.ReadAllBlogs())
	} else {
		fmt.Println("Something wrong!")
		w.WriteHeader(http.StatusInternalServerError)
	}

	ship := shipData{
		Blogs: models.ReadAllBlogs(),
	}

	templ.ExecuteTemplate(w, "index.html", ship)
}
