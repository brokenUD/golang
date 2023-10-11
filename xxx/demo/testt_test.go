package demotest

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
)

type Student1 struct{
	Age int
	Gender bool
	Name string
}

type Class1 struct{
	Id string
	Students []Student1
}

var(
	s = Student1{11, true, "b"}
	c = Class1{
		Id:	"20230401",
		Students: []Student1{s,s,s},
	}
)

func TestJson(t *testing.T){
	bytes, err := json.Marshal(c)
	if err != nil {
		t.Fail()
	}
	var c2 Class1
	err = json.Unmarshal(bytes, &c2)
	if err != nil {
		t.Fail()
	}
	if !(c.Id==c2.Id && len(c.Students)==len(c2.Students)){
		t.Fail()
	}
	fmt.Println("json ok")
}

func TestSonic(t *testing.T){
	bytes, err := sonic.Marshal(c)
	if err != nil {
		t.Fail()
	}
	var c2 Class1
	err = sonic.Unmarshal(bytes, &c2)
	if err != nil {
		t.Fail()
	}
	if !(c.Id==c2.Id && len(c.Students)==len(c2.Students)){
		t.Fail()
	}
	fmt.Println("sonic ok")
}

func BenchmarkJson(b *testing.B){
	for i:=0;i<b.N;i++{
		bytes, _ := json.Marshal(c)
		var c2 Class1
		_ = json.Unmarshal(bytes, &c2)
	}
}

func BenchmarkSonic(b *testing.B){
	for i:=0;i<b.N;i++{
		bytes, _ := sonic.Marshal(c)
		var c2 Class1
		_ = sonic.Unmarshal(bytes, &c2)
	}
}
