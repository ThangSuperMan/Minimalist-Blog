package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderAddBlogPage(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./src/views/*.html")

	if err != nil {
		// log.Fatal("Error here: ")
		fmt.Println("Error here: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templ.ExecuteTemplate(w, "add_blog.html", nil)
}
