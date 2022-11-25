package main

import (
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestEmployee(t *testing.T) {
	tests := []struct {
		desc string
		ID   int
		Name string
	}{
		{"case 1", 2, "happy"},
	}
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	employee(w, req)
	w.Result()
	if reflect.DeepEqual(, tests) {
		t.Errorf(" want %v", tests[0])
	}

}
