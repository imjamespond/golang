module example.com/test

go 1.15

replace example.com/greetings => ../test-module/greetings

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	github.com/labstack/echo/v4 v4.1.17
)
