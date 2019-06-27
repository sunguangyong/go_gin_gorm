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
var connArgs string
var Mydb  DB


func init() {
	Mydb.MyUser = "root"
	Mydb.Password = "123456"
	Mydb.Host = "127.0.0.1"
	Mydb.Port = 3306
	Mydb.Db = "bookshop"
	connArgs = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",  Mydb.MyUser, Mydb.Password, Mydb.Host, Mydb.Port, Mydb.Db )
	db, err = gorm.Open("mysql", connArgs)
	if err != nil {
		panic(err)
	}

}

func (d *DB) connect()(db *gorm.DB) {
	db, err = gorm.Open("mysql", connArgs)
	if err != nil {
		panic(err)
	}
	return db
}


func (d *DB) Create() {
	if !db.HasTable(d) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserInfo{}).Error; err != nil {
			panic(err)
		}
	}
}

func (d *DB) AddOnePassport(username, password string) (err error){
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

func (d *DB) GetOnePassport(username string) (err error, users []UserInfo){
	db, err = gorm.Open("mysql", connArgs)
	users = make([]UserInfo, 0)
	db = db.Where("user_name = ?", username)
	if err := db.Find(&users).Error; err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(users)
	if err != nil {
		return err, users
	}else {
		return nil, users
	}
}
