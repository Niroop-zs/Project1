package main

import (
	"fmt"
)

func prime_check(n int) bool {
	if n <= 0 || n == 1 {
		return false
	} else {
		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				return false
			} else {
				return true
			}
		}
	}
	return true
}
func main() {
	n := 1
	fmt.Println(prime_check(n))
}
