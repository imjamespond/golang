module test-gin-auth

go 1.15

require (
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19
	github.com/gin-gonic/gin v1.7.1
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect 
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.8
)

replace test-gin-auth/controller => ./controller

replace test-gin-auth/service => ./service

replace test-gin-auth/model => ./model

