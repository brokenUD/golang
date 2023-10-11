package ginn

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"

	redis "github.com/redis/go-redis/v9"
)

type Student struct{
	Name string
	Age int
	Height float32
}

type Request struct {
	StudentId string `json:"student_id"`
}

func checkErr(err error){
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetStudentInfo(studentId string)Student{
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})
	ctx := context.TODO()
	stu := Student{}
	for filed, value := range client.HGetAll(ctx, studentId).Val(){
		switch filed{
		case "Name":
			stu.Name = value
		case "Age":
			age, err := strconv.Atoi(value)
			if err == nil {
				stu.Age = age
			}
		case "Height":
			height, err := strconv.ParseFloat(value, 10)
			if err == nil{
				stu.Height = float32(height)
			}
		}
	}
	return stu
}

func GetName(ctx *gin.Context){
	param := ctx.Query("student_id") //
	if len(param) == 0{
		ctx.String(http.StatusBadRequest, "please indidate student_id")
		return
	}
	stu := GetStudentInfo(param)
	ctx.String(http.StatusOK, stu.Name)
}

func GetAge(ctx *gin.Context){
	param := ctx.PostForm("student_id")  //
	if len(param) == 0{
		ctx.String(http.StatusBadRequest, "please indidate student_id")
		return
	}
	stu := GetStudentInfo(param)
	ctx.String(http.StatusOK, strconv.Itoa(stu.Age))
}

func GetHeight(ctx *gin.Context){
	var param Request
	err := ctx.BindJSON(&param)
	if err != nil {
		ctx.String(http.StatusBadRequest, "please indidate student_id in json")
		return
	}
	stu := GetStudentInfo(param.StudentId)
	ctx.String(http.StatusOK, strconv.FormatFloat(float64(stu.Height), 'f', 1, 64))
}

func GinTest(){
	engine := gin.Default()
	engine.GET("/get_name", GetName)    //	路由
	engine.POST("/get_age", GetAge)     //
	engine.POST("/get_height", GetHeight)     //
	err := engine.Run("0.0.0.0:2345")
	if err != nil {
		panic(err)
	}
}
