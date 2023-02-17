package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
func main() {
	StatDynServer()
}
*/

func StatDynServer() {
	http.HandleFunc("/static", serveStaticWebpage) //endpoint "/static"
	http.HandleFunc("/", serveTimeDynamically)     //endpoint "/"
	http.ListenAndServe(PORT, nil)
}

func serveTimeDynamically(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENDPOINT HIT: /")
	response := fmt.Sprintf("Time is %v", time.Now().String())
	fmt.Fprintln(w, response)
}

func serveStaticWebpage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENDPOINT HIT: /static")
	http.ServeFile(w, r, "res/html/static.html")
}
