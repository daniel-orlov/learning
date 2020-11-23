package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
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
