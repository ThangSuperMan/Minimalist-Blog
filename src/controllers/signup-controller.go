package controllers

import (
	"Blog/src/models"
	"Blog/src/structs"
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
func RenderSignupPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RenderSignupPage")
	templ, err := template.ParseGlob("./src/views/*.html")

	if err != nil {
		fmt.Println("Error when parse glob files here", err)
		return
	}

	templ.ExecuteTemplate(w, "signup.html", nil)
}

func SignupAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SignupAccount")
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	var isExist bool = models.CheckUserExists(username)
	var isSamePassword = validatePassword(password, confirmPassword)

	if !isExist && isSamePassword {
		fmt.Println("Can signup for this account")
		var createAt string = time.Now().String()
		models.AddUser(username, password, createAt)

		// Render tmepalte signup page once again
		templ, err := template.ParseGlob("./src/views/*.html")

		notification := structs.NotificationStateSignup{
			Announcement: "signup_successfully",
		}

		if err != nil {
			fmt.Println("Error when parse glob files here", err)
			return
		}

		templ.ExecuteTemplate(w, "signup.html", notification)
		// Return for avoid render one template twice times
		return
	}

	notificationState := structs.NotificationStateSignup{
		Announcement: "signup_unsuccessfully",
		Message:      "Your username it is existing please choose a another one",
	}

	templ, err := template.ParseGlob("./src/views/*.html")
	if err != nil {
		fmt.Println("Error when parse glob files here", err)
		return
	}

	fmt.Println("Can not signup for this account")
	templ.ExecuteTemplate(w, "signup.html", notificationState)
}
