```
docker run --rm -p 9000:9000 -p 8080:8080 --name minio0 \
  -e "MINIO_ROOT_USER=admin" \
  -e "MINIO_ROOT_PASSWORD=admin123456" \
  -v /mnt/data:/data \
  -v /mnt/config:/root/.minio \
  minio/minio server --console-address ":8080" \
  /data

集群，各文件夹要清空
docker run --rm -p 9000:9000 -p 8080:8080 --name minio0 \
  -e "MINIO_ROOT_USER=admin" \
  -e "MINIO_ROOT_PASSWORD=admin123456" \
  -v /mnt/data:/data \
  -v /mnt/config:/root/.minio \
  minio/minio server --console-address ":8080" http://172.17.0.2/data http://172.17.0.3/data http://172.17.0.4/data http://172.17.0.5/data 
docker run --rm --name minio1 \
  -e "MINIO_ROOT_USER=admin" \
  -e "MINIO_ROOT_PASSWORD=admin123456" \
  -v /mnt/data1:/data \
  -v /mnt/config:/root/.minio \
  minio/minio server --console-address ":8080" http://172.17.0.2/data http://172.17.0.3/data http://172.17.0.4/data http://172.17.0.5/data 
docker run --rm --name minio2 \
  -e "MINIO_ROOT_USER=admin" \
  -e "MINIO_ROOT_PASSWORD=admin123456" \
  -v /mnt/data2:/data \
  -v /mnt/config:/root/.minio \
  minio/minio server --console-address ":8080" http://172.17.0.2/data http://172.17.0.3/data http://172.17.0.4/data http://172.17.0.5/data 
docker run --rm --name minio3 \
  -e "MINIO_ROOT_USER=admin" \
  -e "MINIO_ROOT_PASSWORD=admin123456" \
  -v /mnt/data3:/data \
  -v /mnt/config:/root/.minio \
  minio/minio server --console-address ":8080" http://172.17.0.2/data http://172.17.0.3/data http://172.17.0.4/data http://172.17.0.5/data 

docker run --rm -it --entrypoint=/bin/sh minio/mc

```

# 分布式
