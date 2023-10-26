package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func showMysqlVersion() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/TestDB?charset=utf8mb4")
	db.Ping()
	defer db.Close()

	if err != nil {
		fmt.Println("数据库连接失败！")
		log.Fatalln(err)
	}

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(version)
}

func createTable() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/TestDB")
	db.Ping()
	defer db.Close()

	if err != nil {
		fmt.Println("connect DB error !")
		log.Fatalln(err)
	}

	_, err2 := db.Exec("CREATE TABLE user(id INT NOT NULL , name VARCHAR(20), PRIMARY KEY(ID));")
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("successfully create table")
}

func insertItem() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/TestDB?charset=utf8mb4")
	db.Ping()
	defer db.Close()

	if err != nil {
		fmt.Println("connect DB error !")
		log.Fatalln(err)
	}

	_, err2 := db.Query("INSERT INTO user VALUES(1, 'zhangsan')")
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("successfully insert item")
}

func deleteItem() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/TestDB?charset=utf8mb4")
	db.Ping()
	defer db.Close()

	if err != nil {
		fmt.Println("connect DB error !")
		log.Fatalln(err)
	}

	sql := "DELETE FROM user WHERE id = 1"
	res, err2 := db.Exec(sql)

	if err2 != nil {
		panic(err2.Error())
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("delete item success, statement affected %d rows\n", affectedRows)
}

func alterItem() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/TestDB?charset=utf8mb4")
	db.Ping()
	defer db.Close()

	if err != nil {
		fmt.Println("connect DB error !")
		log.Fatalln(err)
	}

	sql := "update user set name = ? WHERE id = ?"
	res, err2 := db.Exec(sql, "lisi", 1)

	if err2 != nil {
		panic(err2.Error())
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("alter item success, statement affected %d rows\n", affectedRows)
}

func queryItem() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/TestDB?charset=utf8mb4")
	db.Ping()
	defer db.Close()

	if err != nil {
		fmt.Println("connect DB error !")
		log.Fatalln(err)
	}

	var mid int = 1

	result, err2 := db.Query("SELECT * FROM user WHERE id = ?", mid)
	if err2 != nil {
		log.Fatal(err2)
	}

	for result.Next() {

		var id int
		var name string

		err = result.Scan(&id, &name)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("id: %d, name: %s\n", id, name)
	}
}

func dropTable() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/TestDB")
	db.Ping()
	defer db.Close()

	if err != nil {
		fmt.Println("connect DB error !")
		log.Fatalln(err)
	}

	_, err2 := db.Exec("DROP TABLE user;")
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("successfully drop table")
}

func main() {
	showMysqlVersion()
	createTable()
	insertItem()
	queryItem()
	alterItem()
	queryItem()
	deleteItem()
	dropTable()
}
