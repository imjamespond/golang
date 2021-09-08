dir=$(dirname $0)
cd $dir
go build main.go
docker build ./ -t testgo:v0.1
docker tag testgo:v0.1 192.168.1.50:5000/testgo:v0.1
docker push 192.168.1.50:5000/testgo:v0.1
docker image prune -f