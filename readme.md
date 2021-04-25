# DB
``docker run --name some-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql``
### Debeaver
报错: **Public Key Retrieval is not allowed**  
驱动属性设置 "``AllowPublicKeyRetrieval=True``", 测试连接OK
# Test
`` go test -timeout 30s -run ^TestMysql$ test-gin-auth/service ``
# Build
`` go build -o server``

# Test clickhouse
``` 
$ docker run -d --name some-clickhouse-server -p 9000:9000 --ulimit nofile=262144:262144 --volume=/mnt/ch-data:/var/lib/clickhouse yandex/clickhouse-server 

```