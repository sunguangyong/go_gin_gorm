package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"fmt"
)

type User struct {
	ID        int    `gorm:"primary_key"`
	UserName  string `gorm:"type:varchar(20);not null;index:username"`
	PassWord  string `gorm:"type:varchar(20);not null;index:password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}



//type User struct {
//	User_id   int `gorm:"primary_key"` //指定主键并自增
//	Name      string
//	Pwd       string
//	CreatedAt time.Time
//	UpdatedAt time.Time
//}


var db *gorm.DB

func main() {
	var err error
	MyUser := "root"
	Password := "123456"
	Host := "127.0.0.1"
	Port := 3306
	Db := "bookshop"
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",  MyUser,Password, Host, Port, Db )
	db, err = gorm.Open("mysql", connArgs)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if !db.HasTable(&User{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
}