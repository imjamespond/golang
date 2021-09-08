dir=$(dirname $0)
app=testapp2:v0.1
cd $dir
go build main.go
docker build ./ -t $app
docker tag $app 192.168.1.50:5000/$app
docker push 192.168.1.50:5000/$app
docker image prune -f