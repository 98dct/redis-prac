package main

import "fmt"

type operationFunc func(x, y int) int

func (o operationFunc) operate(x, y int) int {
	return o(x, y)
}

var operationMap = make(map[string]operationFunc)

func init() {
	operationMap["+"] = add
	operationMap["-"] = minus
	operationMap["*"] = multiply
	operationMap["/"] = divide
	operationMap["%"] = mod
}

func add(x, y int) int {
	return x + y
}

func minus(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func mod(x, y int) int {
	if y == 0 {
		return 0
	}
	return x % y
}

func divide(x, y int) int {
	if y == 0 {
		return 0
	}
	return x / y
}

func operate(operator string, x, y int) int {
	if f, ok := operationMap[operator]; ok {
		return f.operate(x, y)
	}
	return 0
}

func main() {
	fmt.Println(operate("+", 10, 20))
}
