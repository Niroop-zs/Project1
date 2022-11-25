package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func getDbConnection(driver string, connectionString string) (*sql.DB, error) {
	db, err := sql.Open(driver, connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

//var db *sql.DB
//var err error

func main() {
	db, err := getDbConnection("mysql", "niroop:niroopb07@tcp(127.0.0.1:3306)/empdb")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	n := New(db)
	http.HandleFunc("/handler", n.EmployeeHandler)
	http.HandleFunc("/handlerPost", n.EmployeeHandlerPost)
	log.Println("Starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
