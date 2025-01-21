package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	Id uint
	Name string
	Gender string
	Hobby string
}

func main() {
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&UserInfo{})

	u1 := UserInfo{1, "haha", "man", "basketball"}
	// u2 := UserInfo{2, "hehe", "woman", "socer"}

	db.Create(&u1)
	// db.Create(&u2)

	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "socer")
	fmt.Printf("%#v\n", uu)

	db.Model(&u).Update("hobby", "tennis")

	// db.Delete(&u)
	db.Where(&UserInfo{Name: "haha"}).First(&uu)
	fmt.Printf("%#v\n", uu)
}
