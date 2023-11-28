package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: #{err}")
		return
	}
	name := r.FormValue("name")
	address := r.FormValue("address")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ServUsers")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `Users` (`name`, `address`) VALUES ('%s', '%s')", name, address))
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	fmt.Fprintf(w, "POST request successful\n")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello my friends!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", HelloHandler)

	fmt.Printf("Starting server at port 3306\n")
	if err := http.ListenAndServe(":3306", nil); err != nil {
		log.Fatal(err)
	}

	////////////////////////////////////////////////////////////////////

}
