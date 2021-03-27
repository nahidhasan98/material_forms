package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	//Just a message for ensuring the local server is running
	fmt.Println("Local server is listening on port 9000...")

	http.HandleFunc("/", home)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/forgot_password", forgotpassword)
	http.HandleFunc("/dashboard", dashboard)

	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	//serving file from server to client
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))

	//localhost running on port 9000
	http.ListenAndServe(":9000", nil)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/index.gohtml")
	checkErr(err)

	tmpl.Execute(w, nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/index.gohtml")
	checkErr(err)

	tmpl, err = tmpl.ParseFiles("wpage/register.gohtml")
	checkErr(err)

	tmpl.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/index.gohtml")
	checkErr(err)

	tmpl, err = tmpl.ParseFiles("wpage/login.gohtml")
	checkErr(err)

	tmpl.Execute(w, nil)
}

func forgotpassword(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/index.gohtml")
	checkErr(err)

	tmpl, err = tmpl.ParseFiles("wpage/forgot_password.gohtml")
	checkErr(err)

	tmpl.Execute(w, nil)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/index.gohtml")
	checkErr(err)

	tmpl, err = tmpl.ParseFiles("wpage/dashboard.gohtml")
	checkErr(err)

	tmpl.Execute(w, nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is about page")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is contact page")
}
