package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
*
正则表达式：不仅包含特殊字符，还存在量词等修饰字符，形成了字符+重复数量的控制方式
通配符：仅支持几种特殊字符，例如 “%” “_”等没有量词的概念 例如windows的模糊查询、数据库的sql，不同语言下的通配符的实现可能不同

正则表达式广泛应用于各种编程语言，而且保持了语法的高度一致性
正则表达式包含：元字符和普通字符
元字符：
1. ^: 边界匹配符，代表文本的开始，^出现在方括号中，表示与[]内容相反，[^0-9]表示除了0至9外的其他字符
2. $: 边界匹配符，代表文本的结束
3. .: 匹配任意单个字符
4. \: 代表转义字符
5. ?: 量词，匹配0个或1个字符
6. *：量词，匹配0个或多个字符
7. +：量词，匹配1个或多个字符
8. |：匹配分支，类似或操作，例如 x|y 表示匹配x或者y
9. ( )：两者成对出现，表示分组，例如；(abcd)+, 代表一个或者多个字符串“abcd”的组合
10. { }：两者成对出现，表示量词，用来匹配数量范围，A{3} 匹配三个字符A; A{3,} 匹配三个或更多的字符A; A{3,5} 匹配3到5个字符A
11. [ ]：两者成对出现，表示匹配其中任意字符，[abc]表示匹配a、b、c中的任意一个 [0-9]表示匹配0到9的任意字符
*/
func main() {

	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	//test7()
	test8()
}

// 函数式判断是否匹配
func test1() {

	text := "Golang Programing"
	matched, err := regexp.Match("Golang", []byte(text))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("是否匹配：", matched)
}

func test2() {
	text := "Golang Programing"
	matched, err := regexp.Match("^Golang$", []byte(text))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("是否匹配：", matched)
}

// 检查是否匹配
// 先预编译 在解析
func test3() {

	// 匹配一个只含有英文字母的文本
	reg, _ := regexp.Compile(`^[a-zA-Z]+$`)
	text := "golang11"

	match := reg.MatchString(text)
	fmt.Println(text, "是否只含有英文字母: ", match)
}

// 查找匹配位置
func test4() {
	reg, _ := regexp.Compile(`\p{Han}`)
	text := "Golang的学习"
	fmt.Printf("匹配的位置：%v\n", reg.FindStringIndex(text))        // 只返回第一个匹配的位置
	fmt.Printf("匹配的位置：%v\n", reg.FindAllStringIndex(text, -1)) // 返回所有匹配的位置
	fmt.Printf("匹配的位置：%v\n", reg.FindAllStringIndex(text, 1))  // 返回所有匹配中第一次匹配的位置
}

// 获得匹配的文本
func test5() {

	reg, _ := regexp.Compile(`\p{Han}`)
	text := "Golang的学习"
	fmt.Printf("匹配字符：%v\n", reg.FindString(text))
	fmt.Printf("匹配字符：%v\n", reg.FindAllString(text, -1))
}

// 替换文本内容
func test6() {
	reg := regexp.MustCompile(`(\d{4})(\d{2})(\d{2})`)
	fmt.Println(reg.ReplaceAllString("20240820", "$1-$2-$3"))
}

// 替换文本内容
func test7() {
	reg := regexp.MustCompile(`\w+`)
	text := "i am a student, i have go to school everyday. but all the efforts is valuable."
	// 每获得一个匹配项，自定义函数都将被调用一次
	res := reg.ReplaceAllStringFunc(text, func(matched string) string {
		return strings.ToUpper(matched[0:1]) + matched[1:]
	})
	fmt.Println(res)
}

// 在一个连续序列中检查是否存在3,4,5,6,7这样的行为序列
func test8() {

	reg := regexp.MustCompile(`3(.*,4)(.*,5)(.*,6)(.*,7)`)
	fmt.Println("是否匹配:", reg.MatchString("2,3,4,3,3,5,6,3,7,5,6,5,8"))
}
