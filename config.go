package main

import (
	"utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	db_url  = "test:3306"
	db_name = "testdb"
)

var (
	_db *gorm.DB
)

func init() {
	dns := "root:123456@tcp(" + db_url + ")/" + db_name + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	utils.PanicIf(err)

	_db = db
}

func GetDB() *gorm.DB {
	return _db.Debug()
}
