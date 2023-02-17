package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Page struct {
	Title      string
	RawContent string
	Content    template.HTML
	Date       time.Time
	GUID       string
}

func (p Page) TruncatedText() template.HTML {
	chars := 0
	for i, _ := range p.Content {
		chars++
		if chars > 150 {
			return p.Content[:i] + ` ...`
		}
	}
	return p.Content
}

/*
func main() {
	ConnectToDB()
	templates()
}
*/

func templates() {
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

	routes := mux.NewRouter()
	routes.HandleFunc("/page/{guid:[0-9a-zA-Z\\-]+}", ServePageByGUID)
	routes.HandleFunc("/", RedirIndex)
	routes.HandleFunc("/home", ServeIndex)
	http.Handle("/", routes)
	http.ListenAndServe(":8080", nil)
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	var Pages = []Page{}

	//QUERY
	sqlString := "SELECT page_title, page_content, page_date, page_guid FROM pages ORDER BY $1 DESC"
	orderBy := "page_date"
	pages, err := Db.Query(context.Background(), sqlString, orderBy)
	if err != nil {
		err = fmt.Errorf("unable to get index: %w", err)
		_, _ = fmt.Fprint(os.Stderr, err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	defer pages.Close()

	//POPULATING []Pages
	for pages.Next() {
		page := Page{}
		pages.Scan(&page.Title, &page.RawContent, &page.Date, &page.GUID)
		page.Content = template.HTML(page.RawContent)
		Pages = append(Pages, page)
	}
	templateFile := "./templates/index.html"
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		err = fmt.Errorf("unable to parse template %v: %w", templateFile, err)
		_, _ = fmt.Fprint(os.Stderr, err)
	}
	t.Execute(w, Pages)
}

func RedirIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func ServePageByGUID(w http.ResponseWriter, r *http.Request) {
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
	page.Content = template.HTML(page.RawContent)

	//RESPONSE
	templateFile := "./templates/blog.html"
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		err = fmt.Errorf("unable to parse template %v: %w", templateFile, err)
		_, _ = fmt.Fprint(os.Stderr, err)
	}
	t.Execute(w, page)
}
