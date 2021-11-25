package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	engine := gin.Default()

	connStr := "root:111111@tcp(127.0.0.1:3306)/ginsql"

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = db.Exec("CREATE TABLE PERSON(" +
		"ID INT AUTO_INCREMENT PRIMARY KEY," +
		"NAME VARCHAR(12) NOT NULL," +
		"AGE INT DEFAULT 1" +
		");")
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = db.Exec("INSERT INTO PERSON(NAME,AGE)"+
		"VALUE(?,?);", "Lily", 15)
	if err != nil {
		log.Fatal(err)
		return
	} else {
		fmt.Println("数据插入成功")
	}

	rows, err := db.Query("SELECT ID,NAME,AGE FROM PERSON")
	if err != nil {
		log.Fatal(err)
		return
	}

scan:
	if rows.Next() {
		person := new(Person)
		err = rows.Scan(&person.Id, &person.Name, &person.Age)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(person.Id, person.Name, person.Age)
		goto scan
	}

	engine.Run(":8090")
}
