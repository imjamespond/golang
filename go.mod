module sd-2110

go 1.16

require (
	gorm.io/driver/mysql v1.1.2
	gorm.io/driver/sqlserver v1.1.0
	gorm.io/gorm v1.21.16
	my.com/utils v0.0.0-00010101000000-000000000000
)

replace my.com/utils => ../utils
