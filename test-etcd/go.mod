module test-etcd

go 1.16

require (
	codechiev/utils v0.0.0-00010101000000-000000000000
	github.com/go-kit/kit v0.11.0
	github.com/prometheus/client_golang v1.11.0
	go.uber.org/zap v1.18.1 // indirect
	google.golang.org/grpc v1.39.0
)

replace codechiev/utils => ../../utils
