package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"fmt"
)
type DB struct {
	MyUser         string
	Password       string
	Host           string
	Port           int
	Db             string
}

type UserInfo struct {
	ID        int
	UserName  string
	PassWord  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var db *gorm.DB
var err error

func init() {
    dbs := &DB{}
	dbs.MyUser = "root"
	dbs.Password = "123456"
	dbs.Host = "127.0.0.1"
	dbs.Port = 3306
	dbs.Db = "bookshop"
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",  dbs.MyUser, dbs.Password, dbs.Host, dbs.Port, dbs.Db )
	db, err = gorm.Open("mysql", connArgs)
	if err != nil {
		panic(err)
	}
	//dbs.Create()

}


func (d *DB) Create() {
	defer db.Close()
	if !db.HasTable(d) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserInfo{}).Error; err != nil {
			panic(err)
		}
	}
}


func AddOnePassport(username, password string) (err error){
	defer db.Close()
	user := &UserInfo{
		UserName  : username,
		PassWord  : password,
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	}
    fmt.Println(user)
	if err := db.Create(user).Error; err != nil {
		fmt.Printf(err.Error())
		return err
	} else {
		return nil
	}
}

//func (u *UserInfo) GetOnePassport(username, password string) (err error){
//	defer db.Close()
//    db.Model(&UserInfo{}).Where(&UserInfo{UserName:username})
//
//
//}
