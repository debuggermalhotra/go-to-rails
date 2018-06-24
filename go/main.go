package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KloudTrader/htpasswd"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	log.Println("Hello World")
	file := "/tmp/demo.htpasswd"
	name := "joe"
	password := "secret"
	err := htpasswd.SetPassword(file, name, password, htpasswd.HashBCrypt)
	fmt.Println(err)
}
