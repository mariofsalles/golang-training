package main

import (
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página Root!"))
}
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}
func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Fulano!"))
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/home", home)
	http.HandleFunc("/users", users)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
