```
.\bin\windows\zookeeper-server-start.bat .\config\zookeeper.properties
.\bin\windows\kafka-server-start.bat .\config\server.properties

advertised.listeners=PLAINTEXT://192.168.IP:9092 # 代理机向cli告知kafka的地址
```