package routers

import (
	"Blog/src/controllers"
	"fmt"
	"net/http"
)

func HandleHomeRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleHomeRouter")
	controllers.RenderHomePage(w, r)
}

func HandleDetaiBloglRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("router HandleDetailPage")
	controllers.RenderDetailPage(w, r)
}

func HandleAboutRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("router HandleDetailPage")
	controllers.RenderAboutPage(w, r)
}

func HandleAddBlogRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add blog page router")
	controllers.RenderAddBlogPage(w, r)
}

func HandleLoginRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login router")
	controllers.RenderLoginPage(w, r)
}
