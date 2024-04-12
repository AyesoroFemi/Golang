package routes

import (
	"desktop/go-project/controllers"

	"github.com/gorilla/mux"
)


func BookRoutes(router *mux.Router) {
	router.HandleFunc("/getBooks", controllers.GetBooks)
	router.HandleFunc("/getBook/{bookId}", controllers.GetBook)
	router.HandleFunc("/deleteBook/{bookId}", controllers.DeleteBook)
	router.HandleFunc("/addBook", controllers.AddBook)
}
