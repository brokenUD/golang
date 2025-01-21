package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"runtime"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":9001")
	if err != nil {
		fmt.Println(err)
	}

	var name string
	fmt.Print("UserName:")
	_, err = fmt.Scan(&name)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		fmt.Println(file, line, err)
	}

	_, err = conn.Write([]byte(name + "加入聊天室"))
	if err != nil {
		fmt.Println(err)
	}
	//TODO : 读取服务端数据

	go func() {
		var content []byte
		var b [4096]byte

		for {
			n, err := conn.Read(b[:])
			if err != nil {
				fmt.Println(err)
				break
			}
			content = append(content, b[:n]...)
			if n < 4096 {
				fmt.Println(time.Now().Format("2006-01-02 15:04:05"), string(content))
				content = []byte{}
			}
		}
	}() //<--TODO : 读取服务端数据-->
	//TODO : 向服务端发送数据

	for {
		r := bufio.NewReader(os.Stdin)
		b, _, _ := r.ReadLine()
		_, err := conn.Write([]byte(name + ":" + string(b)))
		if err != nil {
			fmt.Println(err)
			break
		}
	} //<--向服务端发送数据-->
}
