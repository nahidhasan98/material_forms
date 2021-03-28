package main

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"html/template"
	"material_forms/model"
	"net/http"
	"net/smtp"
	"strconv"
	"time"
)

func register(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "mysession")
	checkErr(err)

	if r.Method != "POST" {
		if session.Values["isLoggedIn"] == true {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
			//preparing data for sending to frontend
			if session.Values["isLoggedIn"] == nil {
				session.Values["isLoggedIn"] = false
				session.Values["username"] = ""
			}
			data := struct {
				Title      string
				IsLoggedIn bool
				Username   string
			}{
				Title:      "Register | MASTER-ACADEMY",
				IsLoggedIn: session.Values["isLoggedIn"].(bool),
				Username:   session.Values["username"].(string),
			}

			tmpl, err := template.ParseFiles("template/index.gohtml")
			checkErr(err)
			tmpl, err = tmpl.ParseFiles("wpage/register.gohtml")
			checkErr(err)
			tmpl.Execute(w, data)
		}
	} else {
		var uTable model.UserTable
		currentUnixTime := time.Now().Unix()
		token := generateToken()

		r.ParseForm()
		r.Form.Add("bucket", "master_academy")
		r.Form.Add("aid", "user::2")
		r.Form.Add("firstName", "")
		r.Form.Add("lastName", "")
		r.Form.Add("dateOfBirth", "0")
		r.Form.Add("phone", "")
		r.Form.Add("createdAt", strconv.FormatInt(currentUnixTime, 10))
		r.Form.Add("isVerified", "0")
		r.Form.Add("accVerifyToken", token)
		r.Form.Add("accVerifyTokenSentAt", strconv.FormatInt(currentUnixTime, 10))
		r.Form.Add("passResetToken", "")
		r.Form.Add("passResetTokenSentAt", "0")
		r.Form.Add("type", "user")
		r.Form.Add("status", "1")
		insertResponse := db.Insert(r.Form, &uTable)

		if insertResponse.Status == "success" {
			//sending mail verification link to the user mail
			sendMail(r.FormValue("email"), r.FormValue("username"), "http://localhost:9000/verifyemail/"+token)

			fmt.Fprintln(w, "Registration Done")
		} else {
			fmt.Fprintln(w, "Registration Error")
		}
	}
}

func generateToken() string {
	b := make([]byte, 16)
	rand.Read(b)

	hasher := md5.New()
	hasher.Write(b)

	return hex.EncodeToString(hasher.Sum(nil))
}

func sendMail(email, username, link string) {
	auth := smtp.PlainAuth("", "fakenahid@gmail.com", "hqfumidtzssgmdzr", "smtp.gmail.com")
	to := []string{email}

	//var msg []byte
	var body bytes.Buffer

	subject := "Master-Academy Account Verification"
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := fmt.Sprintf("From: Master-Academy \nSubject: %s \nTo:%s \n%s\n\n", subject, email, mimeHeaders)
	body.Write([]byte(msg))

	data := struct {
		Username, Link string
	}{
		Username: username,
		Link:     link,
	}
	tmpl, err := template.ParseFiles("template/mail.gohtml")
	checkErr(err)
	tmpl.Execute(&body, data)

	err = smtp.SendMail("smtp.gmail.com:587", auth, "", to, body.Bytes())
	checkErr(err)
}
