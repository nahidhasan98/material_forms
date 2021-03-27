package main

import (
	"fmt"
	"html/template"
	"material_forms/model"
	"net/http"
	"os"
	"time"

	"github.com/mateors/mcb"
)

var db *mcb.DB

func init() {
	//couchbase connection block
	db = mcb.Connect("localhost", "root", "bootcamp")
	res, err := db.Ping()
	if err != nil {
		fmt.Println(res)
		os.Exit(1)
	}
	fmt.Println(res, err)
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
	tmpl, err := template.ParseFiles("template/index.gohtml")
	checkErr(err)

	//preparing data for sending frontend
	data := struct {
		Title string
	}{
		Title: "Material Forms | MASTER-ACADEMY",
	}

	tmpl.Execute(w, data)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		tmpl, err := template.ParseFiles("template/index.gohtml")
		checkErr(err)

		tmpl, err = tmpl.ParseFiles("wpage/register.gohtml")
		checkErr(err)

		//preparing data for sending frontend
		data := struct {
			Title string
		}{
			Title: "Register | MASTER-ACADEMY",
		}

		tmpl.Execute(w, data)
	} else {
		uTable := model.UserTable{
			FirstName:            "",
			LastName:             "",
			DateOfBirth:          0,
			Phone:                "",
			CreatedAt:            time.Now().Unix(),
			IsVerified:           false,
			AccVerifyToken:       model.GenerateToken(),
			AccVerifyTokenSentAt: time.Now().Unix(),
			PassResetToken:       "",
			PassResetTokenSentAt: 0,
			Type:                 "user",
			Status:               1,
		}

		r.ParseForm()
		r.Form.Add("bucket", "master_academy")
		r.Form.Add("aid", "request::1") //we will update later
		insertResponse := db.Insert(r.Form, &uTable)

		if insertResponse.Status == "success" {
			fmt.Fprintln(w, "Registration Done")
		} else {
			fmt.Fprintln(w, "Registration Error")
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/index.gohtml")
	checkErr(err)

	tmpl, err = tmpl.ParseFiles("wpage/login.gohtml")
	checkErr(err)

	//preparing data for sending frontend
	data := struct {
		Title string
	}{
		Title: "Login | MASTER-ACADEMY",
	}

	tmpl.Execute(w, data)
}

func forgotpassword(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/index.gohtml")
	checkErr(err)

	tmpl, err = tmpl.ParseFiles("wpage/forgot_password.gohtml")
	checkErr(err)

	//preparing data for sending frontend
	data := struct {
		Title string
	}{
		Title: "Request for password reset | MASTER-ACADEMY",
	}

	tmpl.Execute(w, data)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/index.gohtml")
	checkErr(err)

	tmpl, err = tmpl.ParseFiles("wpage/dashboard.gohtml")
	checkErr(err)

	//preparing data for sending frontend
	data := struct {
		Title string
	}{
		Title: "Dashboard | MASTER-ACADEMY",
	}

	tmpl.Execute(w, data)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is about page")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is contact page")
}
