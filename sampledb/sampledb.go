package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type employee1 struct {
	empid int
	name  string
	id    string
	phnno int
	dept  department
}
type department struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open("mysql",
		"niroop:niroopb07@tcp(127.0.0.1:3306)/a2")
	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err)

		return
	}

}

func getEmployees(db *sql.DB) ([]employee, error) {
	rows, err := db.Query("select * from employee inner join dept on employee.id = dept.id")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var employees []employee1

	// 10 records
	for rows.Next() {
		var e employee1
		err = rows.Scan(&e.id, &e.name, &e.empid, &e.dept.id, &e.phnno, &e.dept.name)
		if err != nil {
			return nil, err
		}

		employees = append(employees, e)
	}

	// to be discussed finally
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func getEmployee(db *sql.DB, id int) (employee1, error) {
	var e employee1
	row := db.QueryRow("SELECT * from employee WHERE id=?", id)

	err := row.Scan(&e.empid, &e.name, &e.id, &e.phnno)
	if err != nil {
		return employee1{}, err
	}

	return e, nil
}
