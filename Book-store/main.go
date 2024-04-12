package main

import (
	"desktop/go-project/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routes.BookRoutes(router)
	routes.UserRoutes(router)
	routes.OrderRoutes(router)

	http.ListenAndServe(":8000", router)

}


