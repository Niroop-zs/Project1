package main

import (
	"testing"
)

func TestRoman(t *testing.T) {

	tests := []struct {
		desc    string
		roman   string
		expResp int
	}{
		{"successful conversation from roman to integer for 1", "I", 1},
		{"successful conversation from roman to integer for 5", "V", 5},
		{"successful conversation from roman to integer for 10 ", "X", 10},
		{"successful conversation from roman to integer for 50", "L", 50},
		{"successful conversation from roman to integer for 100", "C", 100},
		{"successful conversation from roman to integer for 500", "D", 500},
		{"successful conversation from roman to integer for 1000", "M", 1000},
		{"Roman to int for like 9,90,4..", "XC", 90},
		{"Invalid Roman numeral", "IC", 0},
		{"Not a roman character", "B", 0},
		{"long number", "CMIV", 904},
		{"Invalid case/input", "", 0},
	}

	for i, tc := range tests {

		resp := roman(tc.expResp)

		if resp != tc.expResp {
			t.Errorf("Test(%v):[%v]: Exp:(%v), Got:(%v)", i, tc.desc, tc.expResp, resp)
		}

	}

}
