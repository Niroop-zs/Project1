package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestEmpHandlerGet(t *testing.T) {
	tests := []struct {
		desc    string
		expCode int
		expResp string
	}{
		{"Get Request", http.StatusOK,

			`[{"id":"1","name":"jay","age":"24","address":"Bengaluru"}]`,
		},
		{
			"Request with unwanted body", http.StatusOK,
			`[{"id":"2","name":"jay","age":"25","address":"Delhi"}]`,
		},
	}
	req, err := http.NewRequest(http.MethodGet, "/handler", nil)
	if err != nil {
		t.Error()
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(EmpHandlerGet)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != tests[0].expCode {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() != tests[0].expResp {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), tests[0].expResp)
	}
}

func TestEmpHandlerPost(t *testing.T) {
	type testcase struct {
		input       empDetail
		expected    interface{}
		statusCode  int
		method      string
		description string
	}

	var testcases = []testcase{

		{
			empDetail{
				"96",
				"Mohit Bajaj",
				"24",
				"Blr",
			},
			"Added successfully",
			http.StatusCreated,
			http.MethodPost,
			"POST Request -Data will be added",
		},

		{
			empDetail{},
			"Method not allowed",
			http.StatusMethodNotAllowed,
			http.MethodPut,
			"Only POST and GET is allowed",
		},
	}

	var employees []empDetail
	handler := EmpHandlerPost
	for i := range testcases {
		input, _ := json.Marshal(testcases[i].input)
		req := httptest.NewRequest(testcases[i].method, "http://localhost:8080/employee", bytes.NewBuffer(input))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		handler(w, req)
		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)
		switch testcases[i].method {
		case http.MethodGet:
			err := json.Unmarshal(body, &employees)
			if err != nil {
				return
			}
			if !reflect.DeepEqual(employees, testcases[i].expected) {
				t.Errorf("Testcase: %d\n"+
					"Input: %v\n"+
					"Expected Output: %v\n"+
					"Actual Output: %v\n"+
					"Description: %v", i+1, testcases[i].input, testcases[i].expected, employees, testcases[i].description)
			}
		default:
			if testcases[i].expected == resp {
				t.Errorf("Testcase: %d\n"+
					"Input: %v\n"+
					"Expected Output: %v\n"+
					"Description: %v", i+1, testcases[i].input, testcases[i].expected, testcases[i].description)
			}
		}
	}
}
