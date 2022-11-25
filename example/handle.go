package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Employee struct {
	ID   int
	Name string
}
func New1(id int,name string)Employee{
	return Employee{ID: id,Name: name}
}
type store struct {
	db *sql.DB
}
func New(d *sql.DB)store{
	return store{
		db: d,
	}
}
func dbConn{
	db, err := sql.Open("mysq","niroop:niroopb07@tcp(127.0.0.1:3306)/a2")
	if err != nil{
		log.Println(err)
	}
}


func main() {
	http.HandleFunc("/", employee)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func employee(w http.ResponseWriter, r *http.Request) {
	e := Employee{ID: 1, Name: "happy"}
	body, err := json.Marshal(e)
	if err != nil {
		b, _ := json.Marshal(err)
		w.Write(b)

		return
	}

	w.Write(body)
}
