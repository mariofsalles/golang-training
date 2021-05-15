package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type Users struct {
	Name  string
	Email string
}

func user(w http.ResponseWriter, r *http.Request) {
	user := Users{"Fulano", "fulano@fulano.com"}
	templates = template.Must(template.ParseGlob("*.html"))
	templates.ExecuteTemplate(w, "home.html", user)
}

func main() {
	http.HandleFunc("/users", user)
	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
