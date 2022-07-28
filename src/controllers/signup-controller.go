package controllers

import (
	"Blog/src/models"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// Local functions
func validatePassword(password string, confirmPassword string) bool {
	if password == confirmPassword {
		return true
	}

	return false
}

// Global functions
func SignupAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SignupAccount")
	// Get info account
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	var isEsxist bool = models.CheckUserExists(username)
	var isSamePassword = validatePassword(password, confirmPassword)

	if !isEsxist && isSamePassword {
		fmt.Println("Can signup for this account")
		var createAt string = time.Now().String()
		models.AddUser(username, password, createAt)
		return
	}

	fmt.Println("Can not signup for this account")

	// http.Redirect(w, r, "/login"+"?sucessfully=true", http.StatusFound)
}

func RenderSignupPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("render signuppage")

	templ, err := template.ParseGlob("./src/views/*.html")

	if err != nil {
		fmt.Println("Error here", err)
		return
	}

	templ.ExecuteTemplate(w, "signup.html", nil)

}
