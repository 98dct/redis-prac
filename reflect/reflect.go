package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func test1() {

	var a = 1.2
	var user interface{}
	user = User{
		Id:   1,
		Name: "dct",
		Age:  25,
	}

	user1 := User{
		Id:   2,
		Name: "zsf",
		Age:  99,
	}

	var sli = []string{"1a", "2b", "3c"}

	// reflect.Type是类型，类型是无限的，包含自定义类型 main.User等
	// reflect.Kind的分类，分类是有限的，可以枚举的，例如 struct
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(&a))
	fmt.Println(reflect.TypeOf(a).Kind())  // float64
	fmt.Println(reflect.TypeOf(&a).Kind()) // ptr
	fmt.Println(reflect.TypeOf(user))
	fmt.Println(reflect.TypeOf(&user))
	fmt.Println(reflect.TypeOf(user).Kind())
	fmt.Println(reflect.TypeOf(&user).Kind())
	fmt.Println(reflect.TypeOf(user1))
	fmt.Println(reflect.TypeOf(&user1))
	fmt.Println(reflect.TypeOf(user1).Kind())
	fmt.Println(reflect.TypeOf(&user1).Kind())

	fmt.Println("-----------------------------")
	// elem只适合 array、slice、map、chan、pointer
	//fmt.Println(reflect.TypeOf(a).Elem())
	//fmt.Println(reflect.TypeOf(user).Elem())
	fmt.Println(reflect.TypeOf(sli).Elem())
	fmt.Println(reflect.TypeOf(&sli).Elem())
	fmt.Println(reflect.TypeOf(&a).Elem())
	fmt.Println(reflect.TypeOf(&user).Elem())
	fmt.Println(reflect.TypeOf(&user1).Elem())

	fmt.Println("-----------------------------")

	// value的elem必须是interface{} 或者 pointer
	fmt.Println(reflect.ValueOf(user))
	fmt.Println(reflect.ValueOf(&user))
	fmt.Println(reflect.ValueOf(&a).Elem())
	fmt.Println(reflect.ValueOf(&user).Elem())
	fmt.Println(reflect.ValueOf(&user1).Elem())
	fmt.Println(reflect.ValueOf(&sli).Elem())
}

func main() {
	test1()
}
