//create: 2016/05/05 15:58:20 change: 2016/05/05 17:06:09 author:lijiao
package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type UserX struct {
	Name    string
	Updated time.Time `xorm:"updated"`
}

func main() {

	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	err = engine.Sync2(new(User))
	if err != nil {
		log.Fatal(err)
	}

	userA := User{
		Name: "userA",
		Age:  13,
	}

	affectd, err := engine.Insert(userA)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Affetced: %d\n", affectd)

	userB := User{
		Name: "userA",
	}

	has, err := engine.Get(&userB)
	if err != nil {
		log.Fatal(err)
	}
	if has {
		fmt.Printf("%v\n", userB)
	} else {
		fmt.Println("Not Found")
	}

	userC := UserX{
		Name: "userA",
	}

	has, err = engine.Get(&userC)
	if err != nil {
		log.Fatal(err)
	}
	if has {
		fmt.Printf("%v\n", userC)
	} else {
		fmt.Println("Not Found")
	}
}
