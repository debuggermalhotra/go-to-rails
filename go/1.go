package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/KloudTrader/htpasswd"
)

func main() {
	http.HandleFunc("/change-password", change_password)
	http.HandleFunc("/password-changed", changed_password)
	server := http.ListenAndServe(":5000", nil)
	if server != nil {
		log.Fatal("ListenAndServe: ", server)
	}
}

func change_password(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("Name:", r.Form["name"])
		fmt.Println("New password:", r.Form["new_password"])

	}
}

func changed_password(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)

	file := "/tmp/demo.htpasswd"
	name := strings.Join(r.Form["name"], " ")
	new_password := strings.Join(r.Form["new_password"], " ")
	err := htpasswd.SetPassword(file, name, new_password, htpasswd.HashBCrypt)

	p, err := htpasswd.ParseHtpasswdFile(file)

	// write data to response
	fmt.Fprintf(w, "Congrats! Your password was successfully changed")

}
