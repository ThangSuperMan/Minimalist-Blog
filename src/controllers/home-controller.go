package controllers

import (
	"Blog/src/models"
	"fmt"
	"html/template"
	"net/http"
	"reflect"
)

// type Ship struct {
//     id int
//     user
// }

func RenderHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("render homepage")
	t, err := template.ParseFiles("./src/views/index.html")

	if err != nil {
		fmt.Println("Error here", err)
		return
	}

	fmt.Println(models.ReadAllBlogs())
	fmt.Println(reflect.TypeOf(models.ReadAllBlogs()))

	if models.ReadAllBlogs() != nil {
		fmt.Println("blogs: ", models.ReadAllBlogs())
	} else {
		fmt.Println("Something wrong!")
		w.WriteHeader(http.StatusInternalServerError)
	}

	// var food [2]string
	// food[0] = "apple"
	// food[1] = "banana"

	// var Data string = "My data :)"

	// student := Student{
	// 	Name: "thang",
	// 	Age:  20,
	// }

	// fmt.Println(food)

	// type Ship struct {
	// 	id    int
	// 	Blogs any
	// }

	// var data = Ship{
	// 	id:    1,
	// 	Blogs: models.ReadAllBlogs(),
	// }

	type bigdata struct {
		SOMETHING string
	}

	data := bigdata{
		SOMETHING: "something good just happened",
	}

	t.Execute(w, data)
}
