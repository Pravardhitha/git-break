package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

const (
	PORT = 9994
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))

	navs := []string{"check", "defend", "about"}

	tpl.ExecuteTemplate(w, "check", navs)
}

func exploitHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))

	navs := []string{"check", "defend", "about"}

	tpl.ExecuteTemplate(w, "check", navs)
}

func defendHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))

	navs := []string{"check", "defend", "about"}

	tpl.ExecuteTemplate(w, "defend", navs)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tpl := template.Must(template.ParseGlob("templates/*.html"))
		navs := []string{"check", "defend", "about"}
		tpl.ExecuteTemplate(w, "about", navs)
	default:
		fmt.Fprintf(w, "err: resource not found")
	}
}

func main() {
	port := strconv.Itoa(PORT)

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/check", exploitHandler)
	http.HandleFunc("/defend", defendHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Printf("http://localhost:%s\n", port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
