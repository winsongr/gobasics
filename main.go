package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@123@/godb"), &gorm.Config{})
	if err != nil {
		panic("db error")
	} else {
		fmt.Println("7 is odd")
	}

	db.AutoMigrate(&user{})
	User := user{
		Name:  "wibi",
		Lastn: "1608",
		Cell:  1608,
		Pin:   1608,
	}
	db.Create(&User)
}

type user struct {
	Id    int
	Name  string
	Lastn string
	Cell  int
	Pin   int
}
