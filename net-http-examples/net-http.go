package main

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	x := r.URL.Path[1:]
	data := "1"

	if len(x) > 0 {
		data = "Merhaba " + x
	} else {
		data = "index page"
	}
	w.Write([]byte(data))
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
