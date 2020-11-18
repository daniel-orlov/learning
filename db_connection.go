package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
	"os"
)

/*
//if we were to parse this into URL-string in runtime
const (
	DBHOST = "127.0.0.1"
	DBPORT = ":5432"
	DBUSER = ""
	DBPASS = ""
	DBDBASE = ""
)
*/

//keeping db-connection global for DRYness
var db *pgxpool.Pool

type Page struct {
	Title   string
	Content string
	Date    string
}

func main() {
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		err = fmt.Errorf("unable to parse DATABASE_URL: %w", err)
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	db, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		err = fmt.Errorf("unable to create connection pool: %w", err)
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	routes := mux.NewRouter()
	routes.HandleFunc("/page/{id:[0-9]+}", ServePage)
	http.Handle("/", routes)
	http.ListenAndServe(":8080", nil)

}

func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	page := Page{}
	fmt.Println(pageID)

	//QUERY
	sqlString := "SELECT page_title, page_content, page_date FROM pages WHERE id=$1"
	err := db.QueryRow(context.Background(), sqlString, pageID).Scan(&page.Title, &page.Content, &page.Date)
	if err != nil {
		err = fmt.Errorf("unable to get page %v: %w", pageID, err)
		_, _ = fmt.Fprint(os.Stderr, err)
	}

	//RESPONSE
	html := fmt.Sprintf("<html><head><title>%v</title></head><body><h1>%v</h1><div>%v</div><div>%v</div></body></html>",
		page.Title, page.Title, page.Content, page.Date,
	)
	_, _ = fmt.Fprintln(w, html)
}
