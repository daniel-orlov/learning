package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

/*
func main() {
	ConnectToDB()
	restApi()
}
*/

func restApi() {
	routes := mux.NewRouter()
	routes.HandleFunc("/api/pages", APIPage).
		Methods("GET").
		Schemes("https")
	routes.HandleFunc("/api/pages/{guid:[0-9a-zA\\-]+}", APIPage).
		Methods("GET").
		Schemes("https")
	routes.HandleFunc("/page/{guid:[0-9a-zA-Z\\-]+}", ServePageByGUID)
	http.Handle("/", routes)
	http.ListenAndServe(PORT, nil)
}

func APIPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	page := Page{}
	fmt.Println(pageGUID)

	//QUERY
	sqlString := "SELECT page_title, page_content, page_date FROM pages WHERE page_guid=$1"
	err := Db.QueryRow(context.Background(), sqlString, pageGUID).Scan(&page.Title, &page.RawContent, &page.Date)
	if err != nil {
		err = fmt.Errorf("unable to get page %v: %w", pageGUID, err)
		_, _ = fmt.Fprint(os.Stderr, err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	//RESPONSE
	APIOutput, err := json.Marshal(page)
	fmt.Printf("%s\n", APIOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		err = fmt.Errorf("unable to marshal %v into json: %w", page, err)
		_, _ = fmt.Fprint(os.Stderr, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, page)
}
