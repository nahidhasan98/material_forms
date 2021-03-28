package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/mateors/mcb"
)

var db *mcb.DB
var store = sessions.NewCookieStore([]byte("mysession"))

func init() {
	//couchbase connection block
	db = mcb.Connect("localhost", "root", "bootcamp")
	res, err := db.Ping()
	if err != nil {
		fmt.Println(res)
		os.Exit(1)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	//Just a message for ensuring the local server is running
	fmt.Println("Local server is listening on port 9000...")

	http.HandleFunc("/", home)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/forgot_password", forgotpassword)
	http.HandleFunc("/dashboard", dashboard)

	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	//serving file from server to client
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))

	//localhost running on port 9000
	http.ListenAndServe(":9000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
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
		Title:      "Material Forms | MASTER-ACADEMY",
		IsLoggedIn: session.Values["isLoggedIn"].(bool),
		Username:   session.Values["username"].(string),
	}

	tmpl, err := template.ParseFiles("template/index.gohtml")
	checkErr(err)
	tmpl.Execute(w, data)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is about page")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is contact page")
}
