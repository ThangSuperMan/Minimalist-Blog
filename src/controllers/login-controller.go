package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RenderLoginPage")

	templ, err := template.ParseGlob("./src/views/*.html")

	if err != nil {
		fmt.Println("Error here", err)
		return
	}

	templ.ExecuteTemplate(w, "login.html", nil)
}
