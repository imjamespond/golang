package loadbalancing

import (
	"log"

	"google.golang.org/grpc/resolver"
)

const (
	testScheme      = "test"
	testServiceName = "test.serivce"
)

var (
	ServerAddrs = []string{"localhost:50051", "localhost:50052"}
)

type TestResolver struct {
	target resolver.Target
	conn   resolver.ClientConn
	// addrs  map[string][]string
}

func (r *TestResolver) start() {
	// targetAddrs := r.addrs[r.target.Endpoint] // 取出当前 service对应的 grpc hosts
	targetAddrs := ServerAddrs
	_adds := make([]resolver.Address, len(targetAddrs))
	for i, addr := range targetAddrs {
		_adds[i] = resolver.Address{Addr: addr}
	}
	r.conn.UpdateState(resolver.State{Addresses: _adds})

	log.Println("TestResolver start Endpoint", r.target.Endpoint)
	log.Println("targetAddrs", targetAddrs, "UpdateState Addresses", _adds)
}
func (r *TestResolver) ResolveNow(opts resolver.ResolveNowOptions) {
	log.Println("ResolveNow")
}
func (r *TestResolver) Close() {
	log.Println("Close")
}

// gopath/pkg/mod/google.golang.org/grpc@v1.39.0/resolver/resolver.go
type TestResolverBuilder struct{}

func (*TestResolverBuilder) Build(target resolver.Target, conn resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	log.Println("TestResolverBuilder build", target)
	testReslover := &TestResolver{
		target: target,
		conn:   conn,
		// 注册每个service对应的 grpc hosts 组
		// addrs: map[string][]string{
		// 	testServiceName: ServerAddrs,
		// },
	}
	testReslover.start()
	return testReslover, nil
}
func (*TestResolverBuilder) Scheme() string {
	return testScheme
}
