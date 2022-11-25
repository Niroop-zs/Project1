package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/handler", EmpHandler)
	fmt.Println("starting a server port 8080 ")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		_ = fmt.Errorf("not reachable")
	}
}
