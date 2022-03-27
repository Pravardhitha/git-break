package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	U "github.com/Pravardhitha/git-break/src/util"
	"github.com/joho/godotenv"
)

const (
	PORT = 9994
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))
	navs := []string{"check", "defend", "about"}
	switch r.Method {
	case "GET":
		tpl.ExecuteTemplate(w, "headnbod", navs)
		tpl.ExecuteTemplate(w, "check", nil)
		tpl.ExecuteTemplate(w, "tail", nil)
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

			tpl.ExecuteTemplate(w, "headnbod", navs)
			// tpl.ExecuteTemplate(w, "defend", nil)
			tpl.ExecuteTemplate(w, "tail", nil)

		} else {
			tpl.ExecuteTemplate(w, "headnbod", navs)
			tpl.ExecuteTemplate(w, "err", nil)
			tpl.ExecuteTemplate(w, "tail", nil)
		}
	default:
		fmt.Fprintf(w, "err: resource not found")
	}
}

func defendHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))
	navs := []string{"check", "defend", "about"}
	switch r.Method {
	case "GET":
		tpl.ExecuteTemplate(w, "headnbod", navs)
		tpl.ExecuteTemplate(w, "defend", nil)
		tpl.ExecuteTemplate(w, "tail", nil)
	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Fatal(err.Error())
		}

		tpl.ExecuteTemplate(w, "headnbod", navs)
		// tpl.ExecuteTemplate(w, "defend", nil)
		tpl.ExecuteTemplate(w, "tail", nil)

	default:
		fmt.Fprintf(w, "err: resource not found")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tpl := template.Must(template.ParseGlob("templates/*.html"))
		navs := []string{"check", "defend", "about"}
		tpl.ExecuteTemplate(w, "headnbod", navs)
		tpl.ExecuteTemplate(w, "about", nil)
		tpl.ExecuteTemplate(w, "tail", nil)
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
	http.HandleFunc("/try", func(w http.ResponseWriter, r *http.Request) {
		U.SendSMS("brotendo")
	})

	fmt.Printf("http://localhost:%s\n", port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
