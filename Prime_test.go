package main

import (
	"testing"
)

func TestPrime_or_Not(t *testing.T) {
	var aop bool
	Test := []struct {
		ip         int
		expectedop bool
	}{
		{1, false},
		{1334, false},
		{61, true},
		{97, true},
		{-1, false},
		{-234, false},
		{0, false},
		{9973, true},
	}
	for _, testcase := range Test {
		aop = prime_check(testcase.ip)
		if aop != testcase.expectedop {
			t.Errorf("ip is (%v), expected(%v),got (%v)", testcase.ip, testcase.expectedop, aop)
		}
	}
}
