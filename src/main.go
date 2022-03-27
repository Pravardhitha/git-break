package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	U "github.com/Pravardhitha/git-break/src/util"
)

const (
	PORT = 9994
)

func checkHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))
	navs := []string{"check", "defend", "about"}
	switch r.Method {
	case "GET":
		tpl.ExecuteTemplate(w, "check", navs)
	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Fatal(err.Error())
		}

		input := r.FormValue("input")
		if U.IsGithub(input) {
			// to get the user and repo from values
			specialFmt := fmt.Sprintf("%sSDUIE$FHWUIFNKWF$", input)
			// fmt.Println(specialFmt)
			gitinfo := U.ParseUserRepo(specialFmt)

			data := U.Branches(gitinfo.User, gitinfo.Repo)
			fmt.Println(data[0].Protected)

		} else {
			tpl.ExecuteTemplate(w, "err", navs)
		}
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
