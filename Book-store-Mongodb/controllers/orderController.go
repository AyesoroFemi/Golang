package controllers

import (
	"desktop/go-project/data"
	"desktop/go-project/datatypes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func OrderBook(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("username")

	vars := mux.Vars(r)
	bookIdstr, errbool := vars["bookId"]

	if !errbool {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Book id not provided in url"))
		return
	}

	bookId, err := strconv.Atoi(bookIdstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Book id not provided in url"))
	}

	_, exists := data.Books[bookId]

	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Book is not in the store"))
		return
	}
	data.Orders[username] = append(data.Orders[username], bookId)
	w.Write([]byte("Book ordered successfully"))
	
}
func ListAllOrderedBooks(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("username")

	var listOfBooks datatypes.Books
	orders := data.Orders[username]

	for _, bookId := range orders {
		book := datatypes.Book{
			Id: bookId,
			Name: data.Books[bookId],
		}
		listOfBooks.AddBooksToList(book)
	}
	marshaledData, err := json.Marshal(listOfBooks)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(marshaledData)
}