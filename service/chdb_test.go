package service

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

type Logger struct {
	ID         uint64 `gorm:"primaryKey"`
	Desc       string
	CreateTime time.Time
}

func TestCH(t *testing.T) {
	db := ConnectToClickHouse()
	// Auto Migrate
	db.AutoMigrate(&Logger{})
	// Set table options
	db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&Logger{})

	// Insert
	db.Create(&Logger{Desc: "foobar", CreateTime: time.Now()})

	// Select
	var desc string
	row := db.First(&Logger{Desc: "foobar"}, "desc = ?", "foobar").Select("desc").Row()
	if row != nil {
		row.Scan(&desc)
		fmt.Println(desc)
	}

	// Batch Insert
	for i := 0; i < 100; i++ {
		stri := strconv.Itoa(i)
		batchSize := 1000
		loggers := make([]Logger, batchSize)
		for j := 0; j < batchSize; j++ {
			logger := Logger{Desc: "logger-" + stri + "-" + strconv.Itoa(j), CreateTime: time.Now()}
			loggers[j] = logger
		}
		db.Create(&loggers)
	}

}
