package main

import "fmt"

func main() {
	forRange()
	recursion()
}

func forRange() {

	res := 0
	for i := 1; i <= 50; i++ {
		res += i
	}
	fmt.Println(res)
}

func recursion() {

	i := 1
	res := 0
	res = recur(i, res)
	fmt.Println(res)
}

func recur(i, res int) int {
	if i > 50 {
		return res
	}

	res = recur(i+1, res)
	return res + i
}
