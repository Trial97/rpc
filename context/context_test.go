package context

import "testing"

type client struct{}

func (*client) Call(ctx Context, serviceMethod string, args, reply interface{}) error { return nil }

func TestWithCancel(t *testing.T) {
	ctx := Background()
	c2, _ := WithCancel(ctx)
	c2.Client = new(client)
	if ctx.Client != nil {
		t.Fatal("Expected empty client")
	}
}
