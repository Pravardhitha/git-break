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
	tpl := template.Must(template.ParseGlob("*.html"))
	tpl.ExecuteTemplate(w, "html", nil)
}

func main() {
	port := strconv.Itoa(PORT)

	http.HandleFunc("/", rootHandler)

	fmt.Printf("http://localhost:%s\n", port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
