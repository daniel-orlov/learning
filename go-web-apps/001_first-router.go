package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
func main() {
	FirstRouter()
}
*/

func FirstRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/test", TestHandler)
	http.Handle("/", router)
	fmt.Println("All set up")
}

func TestHandler(w http.ResponseWriter, r *http.Request) {

}
