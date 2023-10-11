package ginn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestGetStudentInfo(t *testing.T){
	id := "学生1"
	stu := GetStudentInfo(id)
	if len(stu.Name) == 0 {
		t.Fail()
	} else {
		fmt.Printf("%+v\n", stu)
	}
}

func TestGetName(t *testing.T){
	resp, err := http.Get("http://127.0.0.1:2345/get_name?student_id=学生2")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if resp == nil {
		fmt.Println("resp nil ")
		t.Fail()
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(string(bytes))
}

func TestGetAge(t *testing.T){
	resp, err := http.PostForm("http://127.0.0.1:2345/get_age", url.Values{"student_id":[]string{"学生1"}})     //url.Values map[string][]string
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if resp == nil {
		fmt.Println("resp nil ")
		t.Fail()
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(string(bytes))
}

func TestGetHeight(t *testing.T){
	client := http.Client{}
	reader := strings.NewReader(`{"student_id":"学生1"}`)

	request, err := http.NewRequest("POST", "http://127.0.0.1:2345/get_height", reader)     //
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("Content-Type", "application/json")

	resp, err :=client.Do(request)

	if resp == nil {
		fmt.Println("resp nil ")
		t.Fail()
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(string(bytes))
	var stu Student
	err = json.Unmarshal(bytes, &stu)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("%+v\n", stu)
}
