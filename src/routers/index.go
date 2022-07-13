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
