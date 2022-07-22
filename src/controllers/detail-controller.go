package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderDetailPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("controller RenderDetailPage")
	t, err := template.ParseFiles("./src/views/detail_blog.tmpl")
	templateFooter, errFooter := template.ParseFiles("./src/views/footer.tmpl")

	if err != nil {
		fmt.Println("Error here", err)
		return
	}

	if errFooter != nil {
		fmt.Println("Error here", err)
		return
	}

	templateFooter.Execute(w, nil)
	t.Execute(w, nil)
}
