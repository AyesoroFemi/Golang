package middleware

import (
	"desktop/go-project/data"
	"fmt"
	"net/http"
)


func ValidateUser(f http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("username")
		password := r.Header.Get("passwprd")

		if data.Users[username] != password || username == "" {
			w.Write([]byte("Failed to authenticate"))
			return
		}
		f(w, r)
	}
}

func ValidateOwner(f http.HandlerFunc)  http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("username")
		password := r.Header.Get("passwprd")
		IfOwner := r.Header.Get("userType")

		if IfOwner != "owner" {
			w.Write([]byte("you are not owner"))
			return
		}

		if data.Users[username] != password || username == "" {
			w.Write([]byte("Failed to authenticate"))
			return
		}
		f(w, r)
	}
}

func TrackNumberOfRequests(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data.NumberOfRequests = data.NumberOfRequests + 1
		fmt.Println("Request Number : ", data.NumberOfRequests)

		f.ServeHTTP(w, r)
	})
}