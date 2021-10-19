```
docker run --name mysql1 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql
docker run --name mssql -e "ACCEPT_EULA=Y" -e "SA_PASSWORD=yourStrong@Password" -p 1433:1433 -d mcr.microsoft.com/mssql/server
```

# [MSSql](https://docs.microsoft.com/en-us/sql/linux/quickstart-install-connect-docker?view=sql-server-ver15&pivots=cs1-bash)
```
docker exec -it mssql bash
/opt/mssql-tools/bin/sqlcmd -S localhost -U SA -P "yourStrong@Password"
```

# Reverse
https://github.com/xxjwxc/gormt  
https://github.com/smallnest/gen  
https://github.com/volatiletech/sqlboiler  
 