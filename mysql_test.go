package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"test-gorm/model"
	"testing"
	"time"
	"utils"

	"gorm.io/gorm"
)

func init() {
	unixnano := time.Now().UnixNano()
	rand.Seed(unixnano)
	fmt.Println(unixnano)
}

func TestMysqlCreate(t *testing.T) {
	db := GetDB()
	db.AutoMigrate(&model.User{})
	now := time.Now()
	user := model.User{Name: "Jinzhu" + strconv.Itoa(rand.Intn(100)), Age: uint8(rand.Intn(100)), Birthday: &now}
	result := db.Create(&user)       // pass pointer of data to Create
	log.Println(user.ID)             // returns inserted data's primary key
	log.Println(result.Error)        // returns error
	log.Println(result.RowsAffected) // returns inserted records count
}

func TestMysqlQuery(t *testing.T) {
	var rs *gorm.DB
	var user *model.User
	var users []model.User
	db := GetDB()

	user = &model.User{}
	rs = db.First(user)
	utils.FatalIf(rs.Error)
	utils.Log(user)

	user = &model.User{}
	rs = db.Last(user)
	utils.FatalIf(rs.Error)
	utils.Log(user)

	db.Offset(2).Limit(2).Find(&users, []int{1, 2, 3})
	utils.Log(users)

	db.First(&users, "id = ?", 1)
	utils.Log(users)

	db.Limit(3).Where("age > ?", 18).Order("age asc").Find(&users)
	utils.Log(users)
}

func TestMysqlDel(t *testing.T) {
	var rs *gorm.DB
	db := GetDB()

	rs = db.Where("age = ?", 87).Delete(model.User{})
	utils.FatalIf(rs.Error)
	log.Println(rs.RowsAffected)
}

func TestMysqlUpdate(t *testing.T) {
	var rs *gorm.DB
	db := GetDB()
	user := model.User{Name: "Jinzhu"}
	rs = db.First(&user)
	utils.FatalIf(rs.Error)
	log.Println(rs.RowsAffected)
	user.UpdatedAt = time.Now()
	db.Save(user)
}
