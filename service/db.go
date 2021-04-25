package service

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var __db__ *gorm.DB

func GetDB() *gorm.DB {
	return __db__
}

func init() {
	__db__ = connect()
}

func connect() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "root:my-secret-pw@tcp(test:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:my-secret-pw@tcp(test:3306)/test?charset=utf8mb4&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                              // default size for string fields
		DisableDatetimePrecision:  true,                                                                             // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                             // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                             // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                            // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
