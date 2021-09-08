dir=$(dirname $0)
cd $dir
go build main.go
docker build ./ -t testgo:v0.1
docker image prune -f