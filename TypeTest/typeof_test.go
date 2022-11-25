package TypeTest

import (
	"reflect"
	"testing"
)

func TestTypeof(t *testing.T) {
	tests := []struct {
		desc      string
		input     interface{}
		expOutput interface{}
	}{
		{"integer type", 10, 100},
		{"float type", 1.2, 12.6},
		{"string type", "world", "helloworld"},
		{"bool type", true, false},
		{"default", []int{1, 2}, "unknown"},
	}
	for i, tc := range tests {
		actualOutput := typeof(tc.input)
		if !reflect.DeepEqual(actualOutput, tc.expOutput) {
			t.Errorf("test: %v dec: %v got: %v expected: %v", i, tc.desc, actualOutput, tc.expOutput)
		}
	}
}
