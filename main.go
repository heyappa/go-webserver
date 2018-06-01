package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := &router{make(map[string]map[string]http.HandlerFunc)}

	r.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "welcome!")
        fmt.Fprintln(w, "GET (about, users/:id, users/:user_id/addr/:addr_id)")
        fmt.Fprintln(w, "POST (users, users/:user_id/addr)")
	})

	r.HandleFunc("GET", "/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "about")
	})

	r.HandleFunc("GET", "/users/:id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "retrieve user")
	})

	r.HandleFunc("GET", "/users/:user_id/addr/:address_id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "retrieve user's address")
	})

	r.HandleFunc("POST", "/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "create user")
	})

	r.HandleFunc("POST", "/users/:user_id/addr", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "create user's address")
	})

	// 8080 포트로 웹 서버 구동
	fmt.Println("starting web server at 8080")
	http.ListenAndServe(":8080", r)
}
