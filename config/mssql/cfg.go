package mssqlCfg

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"my.com/utils"
)

const (
	db_url  = "test2:1433"
	db_name = "testdb"
)

var (
	_db *gorm.DB
)

func init() {
	dsn := "sqlserver://SA:yourStrong@Password@" + db_url + "?database=" + db_name
	_db = Initialize(dsn)
}

func Initialize(dsn string) *gorm.DB {
	// github.com/denisenkom/go-mssqldb
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	utils.PanicIf(err)
	return db
}

func GetDB() *gorm.DB {
	return _db.Debug()
}
