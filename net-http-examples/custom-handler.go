package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	var i user1
	var w user2

	x1 := &messageHandler{"This is first test!"}
	x2 := &messageHandler{"This is second test!"}

	mux := http.NewServeMux()

	mux.Handle("/1", x1)
	mux.Handle("/2", x2)
	mux.Handle("/user1", i)
	mux.Handle("/user2", w)

	http.ListenAndServe(":8080", mux)
}

type user1 int

func (x user1) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr. User1!")
}

type user2 int

func (x user2) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr. User2!")
}

type messageHandler struct {
	message string
}

func (x *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, x.message)
}
