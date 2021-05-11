```shell
# Go 1.16 and above:
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

# Go 1.15 and below:
# Install sqlboiler v4 and the postgresql driver (mysql, mssql, sqlite3 also available)
# NOTE: DO NOT run this inside another Go module (like your project) as it will
# pollute your go.mod with a bunch of stuff you don't want and your binary
# will not get installed.
GO111MODULE=on go get -u -t github.com/volatiletech/sqlboiler/v4
GO111MODULE=on go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql

```
