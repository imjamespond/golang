package helloworld

import "testing"

func TestClient(t *testing.T) {
	RunClient()("Foobar")
}
