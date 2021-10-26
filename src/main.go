package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Content struct {
	Welcoming string
}

func renderWelcome(responseWriter http.ResponseWriter, request *http.Request) {
	htmlTemplate := template.Must(template.ParseFiles("template/welcome.html"))
	htmlTemplate.Execute(responseWriter, Content{Welcoming: "hey hey"})
}

func main() {

	port := "8080"

	http.HandleFunc("/", renderWelcome)

	fmt.Println("Hello")

	fmt.Printf("Running on Port %s \n", port)
	fmt.Printf("http://localhost:%s \n", port)
	http.ListenAndServe(":"+port, nil)
}
