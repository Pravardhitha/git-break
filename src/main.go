package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strconv"
	U "github.com/Pravardhitha/git-break/src/util"
)

const (
	PORT = 9994
)

func checkHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tpl := template.Must(template.ParseGlob("templates/*.html"))
		navs := []string{"check", "defend", "about"}
		tpl.ExecuteTemplate(w, "check", navs)
	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Fatal(err.Error())
		}

		input := r.FormValue("input")
		if util.isGithub(input) {
			// U.GET()
		}
		fmt.Fprintf(w, "%v\t%v", input, reflect.TypeOf(input))

	default:
		fmt.Fprintf(w, "err: resource not found")
	}
}

func defendHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tpl := template.Must(template.ParseGlob("templates/*.html"))
		navs := []string{"check", "defend", "about"}
		tpl.ExecuteTemplate(w, "defend", navs)
	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Fatal(err.Error())
		}

	default:
		fmt.Fprintf(w, "err: resource not found")
	}
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

	http.HandleFunc("/", checkHandler)
	http.HandleFunc("/check", checkHandler)
	http.HandleFunc("/defend", defendHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Printf("http://localhost:%s\n", port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
