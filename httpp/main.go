package main

import (
	"fmt"
	"io"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request){
	fmt.Println("hello")
	fmt.Fprintln(w, "hello")
}

func Login(w http.ResponseWriter, r *http.Request){
	fmt.Println("login")
	fmt.Fprintln(w, "hello")
}

func GetBaidu(){
	res, err := http.Get("http://ww.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

func HeadUrl(){
	urls := []string{
		"http://www.baidu.com",
		"http://www.taobao.com",
		"http://www.google.com",
	}
	for _, url := range urls {
		res, err := http.Head(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res.Status)
	}
}

func main(){
	// http.HandleFunc("/", Hello)
	// http.HandleFunc("/login", Login)
	// err := http.ListenAndServe("localhost:13001", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// GetBaidu()
	HeadUrl()
}

