package controllers

import (
	"Blog/src/models"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func RenderDetailPage(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./src/views/*.html")

	if err != nil {
		fmt.Println("Error database here: ", err)
	}

	vars := mux.Vars(r)
	id := vars["id"]
	blog := models.QueryABlog(id)

	if blog.Title == "" && blog.Content == "" {
		templ.ExecuteTemplate(w, "detail_blog.html", nil)
	} else {
		templ.ExecuteTemplate(w, "detail_blog.html", blog)
	}
}
