package main

import (
	"net/http"
	"fmt"
)

func main(){
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("req, ping")
		w.Write([]byte("pone"))
	})
	http.ListenAndServe(":8080", nil)
}
