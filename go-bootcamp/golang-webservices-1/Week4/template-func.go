package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	ID     int
	Name   string
	Active bool
}

func IsUserIDOdd(u *User) bool {
	return u.ID%2 != 0
}

func main() {
	tmplFuncs := template.FuncMap{
		"OddUserID": IsUserIDOdd,
	}

	tmpl, err := template.
		New("").
		Funcs(tmplFuncs).
		ParseFiles("func.html")
	if err != nil {
		panic(err)
	}

	users := []User{
		User{1, "George", false},
		User{2, "John", false},
		User{3, "Paul", true},
		User{4, "Ringo", true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "func.html",
			struct {
				Users []User
			}{
				users,
			})
		if err != nil {
			panic(err)
		}
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
