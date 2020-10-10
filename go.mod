module test-gin

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jmoiron/sqlx v1.2.0
	// github.com/lib/pq v1.8.0 // indirect
	gorm.io/driver/mysql v1.0.2
	gorm.io/driver/sqlite v1.1.3
	gorm.io/gorm v1.20.2
	just/model v0.0.0-00010101000000-000000000000
)

replace just/model => ./model
