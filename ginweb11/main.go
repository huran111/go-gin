package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	gorm.Model
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:111111@(localhost:3306)/test?charset=utf8mb4")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&UserInfo{})
	//u1:=UserInfo{1,"Hr","难","s"}
	//db.Create(&u1)
	//查询
	var u UserInfo
	db.First(&u)
	fmt.Printf("u:%#v\n", u)

	//更新
	db.Model(&u).Update("name", "sdfsds")
	db.Delete(&u)

}
