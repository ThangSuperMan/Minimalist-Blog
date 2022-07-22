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

func HandleDetailRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("router HandleDetailPage")
	controllers.RenderDetailPage(w, r)
}

func HandleAboutRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("router HandleDetailPage")
	controllers.RenderAboutPage(w, r)
}
