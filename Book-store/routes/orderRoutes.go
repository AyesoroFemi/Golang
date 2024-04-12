package routes

import (
	"desktop/go-project/controllers"

	"github.com/gorilla/mux"
)

func OrderRoutes(router *mux.Router) {
	router.HandleFunc("/orderBook/{bookId}", controllers.OrderBook)
	router.HandleFunc("/listOrderedBooks", controllers.ListAllOrderedBooks)
}