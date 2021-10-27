package loadbalancing

import (
	"log"

	"google.golang.org/grpc/resolver"
)

const (
	testScheme      = "test"
	testServiceName = "test.serivce"
)

type TestResolver struct {
	target resolver.Target
	conn   resolver.ClientConn
	addrs  map[string][]string
}
type TestResolverBuilder struct{}

func (r *TestResolver) start() {
	targetAddrs := r.addrs[r.target.Endpoint]
	_adds := make([]resolver.Address, len(targetAddrs))
	for i, addr := range targetAddrs {
		_adds[i] = resolver.Address{Addr: addr}
	}
	r.conn.UpdateState(resolver.State{Addresses: _adds})
}
func (r *TestResolver) ResolveNow(opts resolver.ResolveNowOptions) {
	log.Println("ResolveNow")
}
func (r *TestResolver) Close() {
	log.Println("Close")
}

func (*TestResolverBuilder) Build(target resolver.Target, conn resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	testReslover := &TestResolver{
		target: target,
		conn:   conn,
		addrs: map[string][]string{
			testServiceName: addrs,
		},
	}
	testReslover.start()
	return testReslover, nil
}

func (*TestResolverBuilder) Scheme() string {
	return testScheme
}
