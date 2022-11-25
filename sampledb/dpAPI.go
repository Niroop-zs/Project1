package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type employee struct {
	ID      string `json:"Id"`
	Name    string `json:"Name"`
	PhoneNo string `json:"Phnno"`
	dept    dept
}

type dept struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

func GetEmployees(db *sql.DB) ([]employee, error) {
	rows, err := db.Query("select * from Employee")
	if err != nil {
		log.Println("")
	}

	defer rows.Close()

	var employees []employee

	for rows.Next() {
		var e employee
		err = rows.Scan(&e.ID, &e.Name, &e.dept.Id, &e.dept.Name, &e.PhoneNo)
		if err != nil {
			return nil, err
		}

		employees = append(employees, e)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return employees, nil
}

//func GetEmployee(db *sql.DB, id string) (employeeDetails, error) {
//	var e employeeDetails
//	row, _ := db.Query("select * from Employee")
//	err := row.Scan(&e.EmpId, &e.Name, &e.d.Id, &e.Phnno)
//	if err != nil {
//		return employeeDetails{}, err
//	}
//	return e, nil
//}

func EmployeeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	db, err := sql.Open("mysql",
		"niroop:niroopb07@tcp(127.0.0.1:3306)/a2")
	if err != nil {
		log.Println(err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
	emp, err := GetEmployee(db, id)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = json.Marshal(emp)
	if err != nil {
		return
	}
}
