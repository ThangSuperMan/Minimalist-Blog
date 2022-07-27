package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderAboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("render about page")

	templ, err := template.ParseGlob("./src/views/*.html")

	if err != nil {
		fmt.Println("Error here", err)
		return
	}

	templ.ExecuteTemplate(w, "about.html", nil)
}
