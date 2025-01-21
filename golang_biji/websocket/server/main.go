package main

import (
	"log"
	"net/http"
	"strings"
	"unsafe"

	"github.com/gorilla/websocket"
)

/*
go get github.com/gorilla/websocket
*/

/*
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
*/
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		/*
			// 字节转字符串
			fmt.Println(string(p))
			fstr := *(*string)(unsafe.Pointer(&p))
			fmt.Println(fstr)
		*/

		/*
			// 字符串转字节
			fmt.Println(*(*[]byte)(unsafe.Pointer(&fstr)))
		*/

		log.Printf("received: %s", p)

		//lower := strings.ToLower(fstr)
		//字符串转大写
		fstr := *(*string)(unsafe.Pointer(&p))
		upperStr := strings.ToUpper(fstr)

		err = conn.WriteMessage(messageType, *(*[]byte)(unsafe.Pointer(&upperStr)))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebsocket)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
