package routes

import (
	"desktop/go-project/controllers"

	"github.com/gorilla/mux"
)


func UserRoutes(router *mux.Router) {
	router.HandleFunc("/login", controllers.Login)
	router.HandleFunc("/signup", controllers.Signup)
}


