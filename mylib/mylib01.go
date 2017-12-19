package mylib

import "fmt"

func Sum(a, b int) int { return a + b }

func sum(a, b int) int { return a + b }

func Sum1(a, b int) {
	var c int
	c = a + b
	fmt.Printf("%d",c)
}
