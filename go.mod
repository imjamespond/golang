module test-gorm

go 1.16

require (
	gorm.io/driver/mysql v1.1.2
	gorm.io/gorm v1.21.13
	utils v0.0.0-00010101000000-000000000000
)

replace utils => ../utils
