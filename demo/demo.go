package main

import (
	"database/sql"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "niroop:niroopb07@tcp(127.0.0.1:3306)/empdb")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	n := New(db)
	http.HandleFunc("/handler", EmployeeHandler)
	log.Println("Starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
