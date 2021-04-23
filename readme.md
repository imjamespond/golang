# DB
``docker run --name some-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql``
### Debeaver
报错: **Public Key Retrieval is not allowed**  
驱动属性设置 "``AllowPublicKeyRetrieval=True``", 测试连接OK
# Test
`` go test -timeout 30s -run ^TestMysql$ test-gin-auth/service ``
## SQL
### user
```
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `passwd` varchar(100) NOT NULL,
  `create_date` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
```