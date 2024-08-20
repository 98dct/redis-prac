package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func test1() {

	s := "abc"
	fmt.Println(len(s), s)
	p := unsafe.Pointer(&s)
	stringHeaderPointer := (*reflect.StringHeader)(p)
	stringHeaderPointer.Len = 10
	fmt.Println(len(s), s)
}

func main() {

	test1()

}
