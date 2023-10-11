package demotest

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DeptEmp struct{
	EmpNo int
	DeptNo string
	FromDate time.Time
	ToDate time.Time
}

func (DeptEmp)TableName()string{
	return "dept_emp"
}

func read(db *gorm.DB, emptNo int)*DeptEmp{
	var users []DeptEmp
	db.Where("emp_no=?", emptNo).Find(&users)
	if len(users) == 0{
		return nil
	}
	return &users[0]
}

func GormTest(){
	dataSourceName := "root:@tcp(127.0.0.1:3306)/employees?charset=utf8&parseTime=True"
	db, err := gorm.Open("mysql", dataSourceName)
	defer db.Close()
	checkError(err)
	demps := read(db, 10010)
	fmt.Println(demps)

}

func checkError(err error){
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
