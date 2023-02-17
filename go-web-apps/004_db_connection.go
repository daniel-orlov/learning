package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
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

// keeping Db-connection global for DRYness
var Db *pgxpool.Pool

/*
type Page struct {
	Title   string
	Content string
	Date    string
}
*/

/*
func main() {
	dbConnection()
}
*/

func ConnectToDB() {
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		err = fmt.Errorf("unable to parse DATABASE_URL: %w", err)
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	Db, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		err = fmt.Errorf("unable to create connection pool: %w", err)
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func dbConnection() {
	routes := mux.NewRouter()
	//routes.HandleFunc("/page/{id:[0-9]+}", ServePageByID)
	routes.HandleFunc("/page/{guid:[0-9a-zA-Z\\-]+}", ServePageByGUID)
	http.Handle("/", routes)
	http.ListenAndServe(":8080", nil)
}

/*
func ServePageByID(w http.ResponseWriter, r *http.Request) {
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

func ServePageByGUID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	page := Page{}
	fmt.Println(pageGUID)

	//QUERY
	sqlString := "SELECT page_title, page_content FROM pages WHERE page_guid=$1"
	err := Db.QueryRow(context.Background(), sqlString, pageGUID).Scan(&page.Title, &page.Content)
	if err != nil {
		err = fmt.Errorf("unable to get page %v: %w", pageGUID, err)
		_, _ = fmt.Fprint(os.Stderr, err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	//RESPONSE
	html := fmt.Sprintf("<html><head><title>%v</title></head><body><h1>%v</h1><div>%v</div><div>%v</div></body></html>",
		page.Title, page.Title, page.Content, page.Date,
	)
	_, _ = fmt.Fprintln(w, html)
}
*/
