package main

import (
	"fmt"
	"io"
	"net/http"
)

// client
func test1() {

	resp, err := http.Get("https://studygolang.com/readings/recent?limit=1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login for : ", r.URL.Query().Get("username"))
	w.Write([]byte("success"))
}

// server
func test2() {
	http.HandleFunc("/user/login", LoginHandler)
	http.ListenAndServe(":8088", nil)
}

func main() {
	//test1()
	test2()
}
