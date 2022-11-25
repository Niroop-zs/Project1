package Triangle

import (
	"testing"
)

func TestTriangle(t *testing.T) {
	var actualoutput string
	var a int
	var b int
	var c int
	tests := []struct {
		description string
		inputA      int
		inputB      int
		inputC      int
		output      string
	}{
		{"case where no sides are equal", 20, 10, 11, "scalene"},
		{"case where all sides are equal", 20, 20, 20, "equilateral"},
		{"case where a and b are equal", 8, 8, 10, "isosceles"},
		{"case where a and c are equal", 8, 10, 8, "isosceles"},
		{"case where b and c are equal", 10, 8, 8, "isosceles"},
		{"case where sum of b and c less than a", 10, 3, 5, "Not a triangle"},
		{"case where sum of a and b less than c", 5, 3, 10, "Not a triangle"},
		{"case where sum of a and c less than b", 1, 13, 5, "Not a triangle"},
		{"case where sum of b and c equal to a", 10, 3, 7, "Not a triangle"},
		{"case where sum of a and b equal to c", 10, 5, 15, "Not a triangle"},
		{"case where sum of a and c equal to b", 10, 15, 5, "Not a triangle"},
		{"case where all sides are zero", 0, 0, 0, "Not a triangle"},
		{"case where all sides are negative", -9, -5, -2, "Not a triangle"},
	}
	for _, testcase := range tests {
		a = testcase.inputA
		b = testcase.inputB
		c = testcase.inputC

		actualoutput = Triangle(a, b, c)
		if actualoutput != testcase.output {
			t.Errorf("failed")
		}
	}
}

func BenchmarkTriangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Triangle(10, 20, 10)
	}
}
