package demotest

import (
	"github.com/bytedance/sonic"
	"encoding/json"
	"fmt"
)

type Student struct{
	Age int
	Gender bool
	Name string
}

type Class struct{
	Id string
	Students []Student
}

func JsonTest(){
	s := Student{10, true, "a"}
	c := Class{
		Id: "20230101",
		Students: []Student{s,s, s},
	}

	bytes, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json序列化失败", err)
		return
	}
	str := string(bytes)
	fmt.Println(str)
	// fmt.Println(bytes)

	var c2 Class
	err = json.Unmarshal(bytes, &c2)
	if err != nil {
		fmt.Println("unmarshal err ", err)
		return
	}
	fmt.Println(c2)
	fmt.Printf("%v\n", c2)
	fmt.Printf("%+v\n", c2)
	fmt.Printf("%#v\n", c2)
}

func SonicTest(){
	s := Student{10, true, "a"}
	c := Class{
		Id: "20230101",
		Students: []Student{s,s, s},
	}

	bytes, err := sonic.Marshal(c)
	if err != nil {
		fmt.Println("sonic序列化失败", err)
		return
	}
	str := string(bytes)
	fmt.Println(str)
	// fmt.Println(bytes)

	var c2 Class
	err = sonic.Unmarshal(bytes, &c2)
	if err != nil {
		fmt.Println("unmarshal err ", err)
		return
	}
	fmt.Println(c2)
	fmt.Printf("%v\n", c2)
	fmt.Printf("%+v\n", c2)
	fmt.Printf("%#v\n", c2)
}
