package main

import (
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"gorm.io/gorm"

	mssqlCfg "sd-2110/config/mssql"
	"sd-2110/model"
)

var GetDB func() *gorm.DB = mssqlCfg.GetDB

func TestMssqlCreate(t *testing.T) {
	db := GetDB()
	db.AutoMigrate(&model.User{})
	now := time.Now()
	user := model.User{Name: "Jinzhu" + strconv.Itoa(rand.Intn(100)), Age: uint8(rand.Intn(100)), Birthday: &now}
	result := db.Create(&user)       // pass pointer of data to Create
	log.Println(user.ID)             // returns inserted data's primary key
	log.Println(result.Error)        // returns error
	log.Println(result.RowsAffected) // returns inserted records count
}
