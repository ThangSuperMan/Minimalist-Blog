package controllers

import (
	"Blog/src/models"
	"fmt"
	"html/template"
	"net/http"
	"reflect"
)

func RenderAboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("render about")

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

	fmt.Println("Prefect", reflect.TypeOf(models.ReadAllBlogs()))

	ship := shipData{
		Blogs: models.ReadAllBlogs(),
	}

	templ.ExecuteTemplate(w, "about.html", ship)
}
