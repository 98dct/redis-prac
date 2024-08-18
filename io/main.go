package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var pwd, _ = os.Getwd()

// 一次性读取bytes长度的文件内容
func test1() {
	// 只读模式打开一个文件

	f, err := os.Open(pwd + "/io/a.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bytes := make([]byte, 1024)
	n, err := f.Read(bytes)
	if err == nil {
		fmt.Printf("共读出%d个字节\n", n)
		fmt.Printf("文件内容：%s", string(bytes))
	} else {
		panic(err)
	}

}

// 循环读取文件内容
func test2() {

	f, err := os.Open(pwd + "/io/a.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var content []byte
	for {
		tp := make([]byte, 10)
		n, err := f.Read(tp)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		content = append(content, tp[0:n]...)
	}

	fmt.Printf("字节数目%d\n", len(content))
	fmt.Printf("文件内容：%s\n", string(content))

}

// 覆盖写
func test3() {
	f, err := os.Create(pwd + "/io/b.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n, err := f.Write([]byte("新写入的内容"))
	if err == nil {
		fmt.Printf("最终写入的字节%d\n", n)
	} else {
		fmt.Println("写入文件失败", err)
	}
}

// 追加写
func test4() {
	f, err := os.OpenFile(pwd+"/io/b.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n, err := f.Write([]byte("\n追加的内容"))
	if err == nil {
		fmt.Printf("最终写入的字节%d\n", n)
	} else {
		fmt.Println("写入文件失败", err)
	}
}

// 一次性读取文件的所有内容
func test5() {
	data, err := os.ReadFile(pwd + "/io/b.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

// bufio读写
// 读到文件末尾会有io.EOF的错误
func test6() {
	f, err := os.Open(pwd + "/io/b.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		fmt.Print(line) //换行符也读取到了
		if err == io.EOF {
			fmt.Println("读取结束")
			return
		}

	}
}

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	test6()
}
