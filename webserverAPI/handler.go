package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type empDetail struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Age     string `json:"age"`
	Address string `json:"address"`
}

var emp = []empDetail{
	{"1", "jay", "24", "Bengaluru"},
}

func EmpHandler(w http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		EmpHandlerGet(w, request)
	} else if request.Method == "POST" {
		EmpHandlerPost(w, request)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func EmpHandlerGet(writer http.ResponseWriter, request *http.Request) {

	response, err := json.Marshal(emp)
	if err != nil {
		log.Println(err)
	}

	var display empDetail

	err = json.Unmarshal(response, &display)
	if err != nil {
		return
	}
	_, err = writer.Write(response)
	if err != nil {
		return
	}
}

func EmpHandlerPost(writer http.ResponseWriter, request *http.Request) {
	var emp1 empDetail
	a, _ := io.ReadAll(request.Body)
	err := json.Unmarshal(a, &emp1)
	if err != nil {
		return
	}

	emp = append(emp, emp1)
	if err != nil {
		return
	}

	_, err = writer.Write(a)
	if err != nil {
		return
	}

}
