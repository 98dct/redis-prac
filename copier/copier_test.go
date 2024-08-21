package main

import (
	"github.com/jinzhu/copier"
	"testing"
)

type User struct {
	Name  string
	Age   int
	Sex   int
	Phone string
	Class int
}

type Student struct {
	Name   string
	Age    int
	Sex    int
	Phone  string
	Class  int
	Number int
	Grade  float64
}

func NormalCopy() {
	s1 := Student{
		Name:   "dct",
		Age:    25,
		Sex:    1,
		Phone:  "15102910825",
		Class:  5,
		Number: 25,
		Grade:  90,
	}

	var u1 User
	u1.Name = s1.Name
	u1.Age = s1.Age
	u1.Sex = s1.Sex
	u1.Phone = s1.Phone
	u1.Class = s1.Class
	//fmt.Println(u1)
}

// 使用了反射，性能较低
func CopierCopy() {
	s1 := Student{
		Name:   "dct",
		Age:    25,
		Sex:    1,
		Phone:  "15102910825",
		Class:  5,
		Number: 25,
		Grade:  90,
	}

	var u1 User
	copier.Copy(&u1, &s1)
	//fmt.Println(u1)
}

func BenchmarkNormalCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormalCopy()
	}
}

func BenchmarkCopierCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CopierCopy()
	}
}
