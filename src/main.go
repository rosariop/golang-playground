package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func renderWelcome(responseWriter http.ResponseWriter, request *http.Request) {
	htmlTemplate := template.Must(template.ParseFiles("template/welcome.html"))
	htmlTemplate.Execute(responseWriter, Content{Welcoming: "hey hey"})
}

func main() {

	//Database
	fmt.Printf("Starting Database connection! \n")

	// Create database handle and check if driver is present
	db, _ := sql.Open("mysql", "127.0.0.1:3306")
	defer db.Close()

	//Connect to db and check version
	version := ""
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Printf("Connected to: %s", version)

	//Webserver
	port := "8080"
	http.HandleFunc("/", renderWelcome)
	fmt.Println("Hello")
	fmt.Printf("Running on Port %s \n", port)
	fmt.Printf("http://localhost:%s \n", port)
	http.ListenAndServe(":"+port, nil)
}
