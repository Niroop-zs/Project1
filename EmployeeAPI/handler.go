package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
)

type Store struct {
	Db *sql.DB
}

func New(d *sql.DB) Store {
	return Store{
		Db: d,
	}
}

func (s Store) EmployeeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":

		// Query the DB to fetch rows
		rows, err := s.Db.Query("select *from employee e inner join dept d on e.dep_id = d.dep_id")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		allEmp := make([]Employee, 0)
		for rows.Next() {
			var emp Employee
			if err := rows.Scan(&emp.ID, &emp.Name, &emp.PhoneNo, &emp.Dept.ID, &emp.Dept.Name); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			allEmp = append(allEmp, emp)
		}

		var reqEmp []Employee

		switch queryPram := r.URL.Query(); true {

		case queryPram.Has("id"):
			for _, emp := range allEmp {
				if emp.ID == queryPram.Get("id") {
					reqEmp = append(reqEmp, emp)
					break
				}
			}
			if reqEmp == nil {
				w.WriteHeader(http.StatusNotFound)
				if _, err := w.Write([]byte(`{"error":"emp id not found"}`)); err != nil {
					log.Println(err)
					return
				}
			}

		case queryPram.Has("deptid"):
			for _, emp := range allEmp {
				if emp.Dept.ID == queryPram.Get("deptid") {
					reqEmp = append(reqEmp, emp)
				}
			}
			if reqEmp == nil {
				w.WriteHeader(http.StatusNotFound)
				if _, err := w.Write([]byte(`{"error":"deptid not found"}`)); err != nil {
					log.Println(err)
				}
			}

		default:
			reqEmp = allEmp
		}

		// Marshal the required employees and send as response
		respBody, err := json.Marshal(reqEmp)
		if err != nil {
			log.Println(err)
			return
		}
		if _, err := w.Write(respBody); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}
}

func (r Store) EmployeeHandlerPost(writer http.ResponseWriter, request *http.Request) {
	var emp Employee
	var Employees []Employee
	writer.Header().Set("Content-Type", "application/json")
	req, err := ioutil.ReadAll(request.Body)
	if err != nil {
		_, err := fmt.Fprintf(writer, "enter data")
		if err != nil {
			return
		}
	}
	err = json.Unmarshal(req, &emp)
	if err != nil {
		return
	}

	query := "INSERT INTO employee(ID, Name, Id, phonno) VALUES (?, ?,?,?)"
	u := uuid.New()
	_, err = r.Db.Exec(query, u, emp.Name, emp.Dept.ID, emp.PhoneNo)
	fmt.Println(err)
	Employees = append(Employees, emp)

	if err != nil {
		http.Error(writer, http.StatusText(500), 500)
		return
	}

	writer.WriteHeader(http.StatusCreated)

	err = json.Unmarshal(req, &emp)
	if err != nil {
		return
	}
}
