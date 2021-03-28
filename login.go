package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "mysession")
	checkErr(err)

	if r.Method != "POST" {
		if session.Values["isLoggedIn"] == true {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
			//preparing info for sending to frontend
			if session.Values["isLoggedIn"] == nil {
				session.Values["isLoggedIn"] = false
				session.Values["username"] = ""
			}
			data := struct {
				Title      string
				IsLoggedIn bool
				Username   string
			}{
				Title:      "Login | MASTER-ACADEMY",
				IsLoggedIn: session.Values["isLoggedIn"].(bool),
				Username:   session.Values["username"].(string),
			}

			tmpl, err := template.ParseFiles("template/index.gohtml")
			checkErr(err)
			tmpl, err = tmpl.ParseFiles("wpage/login.gohtml")
			checkErr(err)
			tmpl.Execute(w, data)
		}
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")

		queryResponse := db.Query("SELECT `password` FROM master_academy WHERE username = '" + username + "'")

		if len(queryResponse.Result) == 0 {
			fmt.Fprintln(w, "Username not found!")
		} else {
			originalPassword := queryResponse.Result[0].(map[string]interface{})["password"]

			if password == originalPassword {
				session.Values["username"] = username
				session.Values["password"] = password
				session.Values["isLoggedIn"] = true
				session.Save(r, w)

				fmt.Fprintln(w, "Login Done")
			} else {
				fmt.Fprintln(w, "Wrong password!")
			}
		}
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "mysession")
	checkErr(err)

	session.Values["username"] = ""
	session.Values["password"] = ""
	session.Values["isLoggedIn"] = false

	session.Options.MaxAge = -1 //cookies will be deleted immediately.
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func forgotpassword(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "mysession")
	checkErr(err)

	//preparing data for sending frontend
	if session.Values["isLoggedIn"] == nil {
		session.Values["isLoggedIn"] = false
		session.Values["username"] = ""
	}
	data := struct {
		Title      string
		IsLoggedIn bool
		Username   string
	}{
		Title:      "Request for password reset | MASTER-ACADEMY",
		IsLoggedIn: session.Values["isLoggedIn"].(bool),
		Username:   session.Values["username"].(string),
	}

	tmpl, err := template.ParseFiles("template/index.gohtml")
	checkErr(err)
	tmpl, err = tmpl.ParseFiles("wpage/forgot_password.gohtml")
	checkErr(err)
	tmpl.Execute(w, data)
}
