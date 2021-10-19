package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"my.com/utils"

	"gorm.io/gorm"

	mysqlCfg "sd-2110/config/mysql"
	"sd-2110/model"
)

var GetMysqlDB func() *gorm.DB = mysqlCfg.GetDB

func init() {
	unixnano := time.Now().UnixNano()
	rand.Seed(unixnano)
	fmt.Println(unixnano)
}

func TestMysqlCreate(t *testing.T) {
	db := GetMysqlDB()
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
	db := GetMysqlDB()

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
	for _, user := range users {
		log.Println(user.Name, user.Age, utils.Date(user.Birthday), utils.Date(&user.CreatedAt), utils.Date(&user.UpdatedAt))
	}
}

func TestMysqlDel(t *testing.T) {
	var rs *gorm.DB
	db := GetMysqlDB()

	rs = db.Where("age = ?", 87).Delete(model.User{})
	utils.FatalIf(rs.Error)
	log.Println(rs.RowsAffected)
}

func TestMysqlUpdate(t *testing.T) {
	var rs *gorm.DB
	db := GetMysqlDB()
	user := model.User{Name: "Jinzhu"}
	rs = db.First(&user)
	utils.FatalIf(rs.Error)
	log.Println(rs.RowsAffected)
	user.UpdatedAt = time.Now()
	db.Save(user)
}

func TestMysqlModel(t *testing.T) {
	var rs *gorm.DB
	db := GetMysqlDB()
	user := model.User{ID: 1}
	// Update single column
	rs = db.Model(&user).Update("UpdatedAt", time.Now())
	utils.FatalIf(rs.Error)
	log.Println(rs.RowsAffected)

	rs = db.Model(&model.User{}).Where("name = ?", "Jinzhu").Update("age", rand.Intn(100))
	utils.FatalIf(rs.Error)
	log.Println(rs.RowsAffected)
}
