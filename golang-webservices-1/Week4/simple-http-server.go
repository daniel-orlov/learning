package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello! Is it me you looking for?")
	w.Write([]byte("!?!"))
}

func main () {
	http.HandleFunc("/", handler)

	fmt.Println("launching server at :8080")
	http.ListenAndServe(":8080", nil)
}
