package main

type Department struct {
	ID   string `json:"ID"`
	Name string `json:"name"`
}

type Employee struct {
	ID      string     `json:"ID"`
	Name    string     `json:"name"`
	PhoneNo string     `json:"phoneNo"`
	Dept    Department `json:"dept"`
}
