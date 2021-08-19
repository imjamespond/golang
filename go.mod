module test-etcd

go 1.16

require (
	go.etcd.io/etcd/api/v3 v3.5.0 // indirect
	go.etcd.io/etcd/client/v3 v3.5.0
	utils v0.0.0-00010101000000-000000000000
)

replace utils => ../utils
