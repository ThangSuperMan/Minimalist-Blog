package controllers

import (
	"Blog/src/models"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
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

	fmt.Println("Username: ", username)
	fmt.Println("Password: ", password)
	fmt.Println("confirmPassword: ", confirmPassword)

	var isEsxist bool = models.CheckUserExists(username)
	fmt.Println("isEsxist: ", isEsxist)
	if isEsxist {
		// Can not sign up for this account
		fmt.Println("Can not sign up for this account")

	} else {
		fmt.Println("You can signup :)")
		var canRegister bool = validatePassword(password, confirmPassword)
		if canRegister {
			fmt.Println("canRegister: ", canRegister)
			// Create table user
			db, err := sql.Open("sqlite3", "test.db")

			if err != nil {
				log.Fatal("Error here: ", err)
			}

			defer db.Close()

			var statement string = `
		CREATE TABLE IF NOT EXISTS  "user" (
		"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username"	TEXT NOT NULL,
		"password"	TEXT
	);`

			result, err := db.Exec(statement)
			if err != nil {
				log.Fatal("Error here: ", err)
			}

			fmt.Println("Result: ", result)

			var query string = "INSERT INTO user (username, password) VALUES (?, ?)"

			stmt, _ := db.Prepare(query)
			resultQuery, _ := stmt.Exec(username, password)
			fmt.Println("ResultQuery", resultQuery)
		}
	}

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
