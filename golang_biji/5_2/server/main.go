package main

import (
	"fmt"
	"net"
	// "net/http"
	"runtime"
)

func main() {
	// http.ListenAndServe(":9001", nil)
	// todo 绑定端口 监听端口
	listen, err := net.Listen("tcp", ":9001")
	if err != nil {
		fmt.Println(err)
	}

	list := make(map[net.Conn]net.Conn, 100)
	ch := make(chan []byte, 1000)

	// todo 阻塞直到客户端链接
	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				_, file, line, _ := runtime.Caller(0)
				fmt.Println(file, line, err)
				break
			}

			list[conn] = conn

			// todo 接受客户端消息 方法1
			go func() {
				for {
					// （头信息+最大允许输入字数） *3
					buf := make([]byte, 1024)
					n, err := conn.Read(buf)
					if err != nil {
						delete(list, conn)
						conn.Close()
						_, file, line, _ := runtime.Caller(0)
						fmt.Println(file, line, err)
						break
					}
					fmt.Println("111", string(buf))
					ch <- buf[:n]
				}
			}() // 接收客户端发送的消息

			// todo ：接受客户端发送的消息 方法2
			go func() {
				var bb [4096]byte
				var content []byte
				for {
					n, err := conn.Read(bb[:])
					if err != nil {
						delete(list, conn)
						conn.Close()
						_, file, line, _ := runtime.Caller(0)
						fmt.Println(file, line, err)
						break
					}

					content = append(content, bb[:n]...)
					fmt.Println("222", string(content))
					if n < 4096 {
						ch <- content
						content = []byte{}
					}
				}
			}() // 接收客户端发送的消息
		}
	}() // 阻塞直到客户端消息

	// todo 向客户端发送消息
	for {
		select {
		case content := <-ch:
			for _, v := range list {
				v.Write(content)
			}
		}
	}
}
