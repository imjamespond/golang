package service

import (
	"fmt"
	"testing"
	"time"

	"test-gin-auth/model"
)

func now() *time.Time {
	now := time.Now()
	return &now
}

func TestMysqlCreate(t *testing.T) {
	db := ConnectMysql()

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	user := model.User{Username: "test003", Passwd: "12345", CreateDate: now()}
	result := db.Create(&user) // pass pointer of data to Create
	// fmt.Println(user)
	if result.Error != nil {
		fmt.Println(result.Error)
	} else {
		fmt.Println(result.RowsAffected)
	}

}

func TestMysqlFind(t *testing.T) {
	db := ConnectMysql()

	var users []model.User
	db.Find(&users)
	fmt.Println(users)
}

func TestMysqlScan(t *testing.T) {
	db := ConnectMysql()

	rows1, err := db.Table("users").Select("*").Rows()
	if err != nil {
		panic(err)
	}
	for rows1.Next() {
		var _user model.User
		db.ScanRows(rows1, &_user)
		fmt.Println(_user)
	}

	rows, err := db.Table("users").Select("username,passwd").Rows()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var name string
		var passwd string
		rows.Scan(&name, &passwd)
		fmt.Printf("name: %v, passwd: %v\n", name, passwd)
	}
}
