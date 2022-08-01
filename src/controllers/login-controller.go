package controllers

import (
	"Blog/src/models"
	// "Blog/src/structs"
	"fmt"
	"html/template"
	"net/http"
)

type UserAuthenticated struct {
	Username string
	IsUser   bool
}

func RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RenderLoginPage")
	templ, err := template.ParseGlob("./src/views/*.html")

	if err != nil {
		fmt.Println("Error here", err)
		return
	}

	// Default render with user not login
	templ.ExecuteTemplate(w, "login.html", nil)
}

func RenderHomePageWithUserAuthentication(w http.ResponseWriter, r *http.Request, username string) {
	fmt.Println("RenderHomePageWithUserAuthentication")
	templ, err := template.ParseGlob("./src/views/*.html")

	if err != nil {
		fmt.Println("Error here", err)
		return
	}

	user := UserAuthenticated{
		Username: username,
		IsUser:   true,
	}

	templ.ExecuteTemplate(w, "index.html", user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginUser Post controllers")
	r.ParseForm()

	usernameInput := r.FormValue("username")
	passwordInput := r.FormValue("password")
	fmt.Println("username: ", usernameInput)
	fmt.Println("password: ", passwordInput)

	// var user structs.User
	user := models.ReadAccountUser(usernameInput)

	fmt.Println("Info user: ", user)
	fmt.Println("Id: ", user.Id)
	fmt.Println("Username: ", user.Username)
	fmt.Println("Username: ", user.Password)

	if user.Id == 0 && user.Username == "" && user.Password == "" {
		fmt.Println("empty user")
	} else {
		// User exists
		if passwordInput == user.Password {
			fmt.Println("Allow user login")
			RenderHomePageWithUserAuthentication(w, r, usernameInput)
		}
	}
}
