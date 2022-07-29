package controllers

import (
	"Blog/src/models"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type NotificationSignupAccount struct {
	Announcement bool
}

func RenderSignupPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RenderSignupPage")
	// sucessfullyJustSignup := r.URL.Query().Get("sucessfully")
	// fmt.Println("sucessfullyJustSignup: ", sucessfullyJustSignup)
	// isAnnouce, _ := strconv.ParseBool(sucessfullyJustSignup)

	notification := NotificationSignupAccount{
		Announcement: false,
	}

	templ, err := template.ParseGlob("./src/views/*.html")

	if err != nil {
		fmt.Println("Error when parse glob files here", err)
		return
	}

	templ.ExecuteTemplate(w, "signup.html", notification)
	// if isAnnouce {
	// 	fmt.Println("Let's annouce")
	// 	templ.ExecuteTemplate(w, "signup.html", annouce)
	// } else {
	// 	fmt.Println("Let's not annouce")
	// 	templ.ExecuteTemplate(w, "signup.html", annouce)
	// }

}

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

		// http.Redirect(w, r, "/signup?successfully=true", http.StatusFound)

		// Render tmepalte signup page once again
		templ, err := template.ParseGlob("./src/views/*.html")

		nontification := NotificationSignupAccount{
			// Announcement: "Congratulations you just successfully signup an anncoue, login and have fun now :)",
			Announcement: true,
		}

		if err != nil {
			fmt.Println("Error when parse glob files here", err)
			return
		}

		templ.ExecuteTemplate(w, "signup.html", nontification)
		// Return for avoid two page show up at the same time
		return
	}

	nontification := NotificationSignupAccount{
		// Announcement: "Congratulations you just successfully signup an anncoue, login and have fun now :)",
		Announcement: false,
	}

	templ, err := template.ParseGlob("./src/views/*.html")
	if err != nil {
		fmt.Println("Error when parse glob files here", err)
		return
	}

	fmt.Println("Can not signup for this account")
	templ.ExecuteTemplate(w, "signup.html", nontification)

	// http.Redirect(w, r, "/signup?successfully=false", http.StatusFound)
}
