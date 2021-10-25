module sd-2110

go 1.16

require (
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/kisielk/gotool v1.0.0 // indirect
	google.golang.org/grpc v1.41.0 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gorm.io/driver/mysql v1.1.2
	gorm.io/driver/sqlserver v1.1.0
	gorm.io/gorm v1.21.16
	my.com/utils v0.0.0-00010101000000-000000000000
)

replace my.com/utils => ../utils
