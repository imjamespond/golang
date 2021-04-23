# DB
``docker run --name some-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql``
### Debeaver
报错: **Public Key Retrieval is not allowed**  
驱动属性设置 "``AllowPublicKeyRetrieval=True``", 测试连接OK
# Test
`` go test -timeout 30s -run ^TestMysql$ test-gin-auth/service ``
