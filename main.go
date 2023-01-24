package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

func main() {
	portNumber := flag.Int("p", 7070, "port number")
	flag.Parse()

	http.HandleFunc("/", GetRoot)

	log.Printf("Serving at http://127.0.0.1:%d", *portNumber)
	err := http.ListenAndServe(":"+strconv.Itoa(*portNumber), nil)

	if err != nil {
		log.Printf("Failed to start server: %s\n", err)
		os.Exit(1)
	}

}

func GetRoot(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "[404] Nothing to see here.", http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("./web/home.html"))

	pageInfo := HomePage{
		MainTitle:     "Air delivery scanner",
		SubTitle:      "This is a demo application that works with fictional data.",
		CopyRightText: "&copy; 2023 | Air delivery scanner project | All rights reserved.",
	}

	tmpl.Execute(w, pageInfo)
}

type HomePage struct {
	MainTitle     string
	SubTitle      string
	CopyRightText string
}
