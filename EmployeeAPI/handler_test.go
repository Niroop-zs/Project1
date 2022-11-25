package main

import (
	"bytes"
	_ "database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	_ "net/http/httptest"
	"testing"
)

var err error
var mock sqlmock.Sqlmock

func TestEmployeeHandler(t *testing.T) {

	Db, mock, err = sqlmock.New()
	if err != nil {
		log.Println(err)
		return
	}
	tests := []struct {
		desc    string
		expCode int
		expResp string
	}{
		{"Get Request", http.StatusOK,

			`[{"ID":"101","name":"raj","depID":"1","phnno":"987654320"}]`,
		},
		//{
		//	"case for unwanted body",http.StatusCreated,
		//	`[{"ID":"102","name":"jay","depID":"1","phnno":"98989777788"}]`,
		//},

	}
	rows := sqlmock.NewRows([]string{"deptId", "depName"})
	rows.AddRow("49682d52-6ab5-44aa-a817-fa4c5b05e086", "luffy")
	mock.ExpectQuery("SELECT (.+) FROM dept\n").WillReturnRows(rows)
	for _, tc := range tests {
		mockReq, _ := http.NewRequest("GET", "/handler", bytes.NewReader(nil))
		mockResp := httptest.NewRecorder()
		EmployeeHandler(mockResp, mockReq)
		assert.Equal(t, tc.expCode, mockResp.Code, tc.desc)
		assert.Equal(t, tc.expResp, mockResp.Body.String())
	}

	rows, err := http.NewRequest(http.MethodGet, "/handler", nil)
	if err != nil {
		t.Error()
	}
	rr := httptest.NewRecorder()
	//establish mock
	EmployeeHandler(rr, req)

	if status := rr.Code; status != tests[0].expCode {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() != tests[0].expResp {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), tests[0].expResp)
	}
}

//	func TestEmployeeHandlerPost{
//		tests := []struct {
//			description string
//			reqBody     Employee
//			mockSqlFunc func(e Employee)
//			expCode     int
//		}{
//			{
//				description: "Post a valid employee",
//				reqBody:     Employee{Name: "Deadpool", PhoneNo: "3213213211", Dept: Department{ID: "", Name: "Software"}},
//				expCode: http.StatusCreated,
//			},
//
// }
func TestEmployeeHandlerPost(t *testing.T) {
	input := Employee{
		"12", "tom", 898989898, "golang",
	}
	Db, mock, err = sqlmock.New()
	reqBody, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/employeePost", bytes.NewBuffer(reqBody))

	w := httptest.NewRecorder()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO emp").WithArgs("122", "jjjj", 9898977775, "java").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	EmployeeHandlerPost(w, req)
	resp := w.Result()
	if resp.StatusCode == http.StatusCreated {
		t.Log("Updated succesfully")
	}
}
