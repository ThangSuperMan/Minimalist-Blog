package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	// "strconv"
)

type AnnouceStateSignup struct {
	isSignupSuccessfully bool
}

func RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RenderLoginPage")
	// sucessfullyJustSignup := r.URL.Query().Get("sucessfully")
	// isAnnouce, _ := strconv.ParseBool(sucessfullyJustSignup)

	// if isAnnouce {
	// 	fmt.Println("Let's annouce")
	// }

	// annouceStateSignup := AnnouceStateSignup{
	// 	isSignupSuccessfully: isAnnouce,
	// }

	templ, err := template.ParseGlob("./src/views/*.html")

	if err != nil {
		fmt.Println("Error here", err)
		return
	}

	templ.ExecuteTemplate(w, "login.html", nil)
}
