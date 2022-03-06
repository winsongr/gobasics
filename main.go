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
	// db.Migrator().DropTable(&user{})
	// db.Migrator().CreateTable(&user{})
	// User := user{
	// 	Name:  "wibi",
	// 	Lastn: "1608",
	// 	Cell:  1608,
	// 	Pin:   1608,
	// }
	// db.Create(&User)

	// User := user{
	// 	Id:    2,
	// Name:  "wibi1608",
	// Lastn: "wibi",
	// Cell:  160868,
	// Pin:   160816,
	// }
	// db.Updates(&User)
	// User := user{
	// 	Id: 4,
	// }
	// db.Delete(&User)
	// 	User := []user{
	// 		{Name: "wibi1608",
	// 			Lastn: "wibi",
	// 			Cell:  160868,
	// 			Pin:   160816},
	// 		{Name: "wibi1608",
	// 			Lastn: "wibi",
	// 			Cell:  160868,
	// 			Pin:   160816},
	// 		{Name: "wibi1608",
	// 			Lastn: "wibi",
	// 			Cell:  160868,
	// 			Pin:   160816},
	// 	}
	// 	for _, User := range User {
	// 		db.Create(&User)
	// 	}
	// User := user{
	// 	Model: gorm.Model{
	// 		CreatedAt: time.Now(),
	// 	},
	// }
	// db.First(&User)
	// db.Last(&User)
	// db.Where("Cell", "1608").First(&User)
	// fmt.Println(&User)
User:=user{

}
db.Create(&User)
}

type user struct {
	gorm.Model
	Id    int
	Name  string
	Lastn string
	Cell  int`gorm:"default:1"`
	Pin   int `gorm:"size:100;null"`
	Email string `gorm:"unique; not null"`
}
