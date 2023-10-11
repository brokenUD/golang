package demotest

import (
	"bufio"
	"fmt"
	"os"
	"io"
	// "strconv"
)

func FileRead(){
	file, err := os.Open("xxx/demo/a.txt")
	if err != nil {
		fmt.Println("open faild ", err)
		return
	}
	defer file.Close()

	content := make([]byte, 100)
	n, err := file.Read(content)
	if err != nil {
		fmt.Println("read faild ", err)
		return
	}
	fmt.Println(string(content[:n]))
	fmt.Println(content[:n])
}


func FileWrite(){
	file, err := os.OpenFile("xxx/demo/b.txt",
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open faild ", err)
		return
	}
	defer file.Close()
	content := "格局"
	n, err := file.Write([]byte(content))
	if err != nil {
		fmt.Println("write faild ", err)
	} else {
		fmt.Println("write succeed ", n)
	}
}

func FileBuffer(){
	file, err := os.Open("xxx/demo/b.txt")
	if err != nil {
		fmt.Println("打开shibai", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF{ // end of file
				break
			} else {
				fmt.Println("du error ", err)
				return
			}
		} else {
			fmt.Print(line)
		}
	}
}
