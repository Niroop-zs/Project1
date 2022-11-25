package ReverseSlice

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		desc      string
		input     []int
		expOutput []int
	}{
		{"Normal case", []int{1, 2, 3}, []int{3, 2, 1}},
		{"same numbers", []int{1, 1, 1}, []int{1, 1, 1}},
		{"Empty slice", make([]int, 0), make([]int, 0)},
	}
	for i, tc := range tests {
		actualOutput := reverse(tc.input)
		if !reflect.DeepEqual(actualOutput, tc.expOutput) {
			t.Errorf("test: %v dec: %v got: %v expected: %v", i, tc.desc, actualOutput, tc.expOutput)
		}
	}
}
