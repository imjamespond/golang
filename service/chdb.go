package service

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

func ConnectToClickHouse() *gorm.DB {
	dsn := "tcp://test:9000?database=testch&read_timeout=10&write_timeout=20"
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(100)

	return db
}
