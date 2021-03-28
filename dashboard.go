package main

import (
	"html/template"
	"net/http"
)

func dashboard(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "mysession")
	checkErr(err)

	if session.Values["isLoggedIn"] == true {
		//preparing data for sending to frontend
		data := struct {
			Title      string
			IsLoggedIn bool
			Username   string
		}{
			Title:      "Dashboard | MASTER-ACADEMY",
			IsLoggedIn: session.Values["isLoggedIn"].(bool),
			Username:   session.Values["username"].(string),
		}

		tmpl, err := template.ParseFiles("template/index.gohtml")
		checkErr(err)
		tmpl, err = tmpl.ParseFiles("wpage/dashboard.gohtml")
		checkErr(err)
		tmpl.Execute(w, data)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
