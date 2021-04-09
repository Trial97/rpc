package rpc

import (
	contextg "context"
	"net"
	"testing"

	"github.com/cgrates/rpc/birpc"
	"github.com/cgrates/rpc/context"
	"github.com/cgrates/rpc/rpcc"
)

func BenchmarkBirpcInArgs(b *testing.B) {
	newServer := birpc.NewBirpcServer()
	newServer.Register(new(birpc.Airth3))
	c1, c2 := net.Pipe()
	go newServer.ServeConn(c1)

	client := birpc.NewBirpcClient(c2)
	defer client.Close()
	client.Register(new(birpc.Airth3))
	ctx := contextg.Background()

	// Synchronous calls
	args := &birpc.Args3{7, 8}
	reply := new(birpc.Reply3)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := client.Call(ctx, "Airth3.Add", args, reply); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBirpcInContext(b *testing.B) {
	newServer := NewBirpcServer()
	newServer.Register(new(Airth2))
	c1, c2 := net.Pipe()
	go newServer.ServeConn(c1)

	client := NewBirpcClient(c2)
	defer client.Close()
	client.Register(new(Airth2))
	ctx := context.Background()

	// Synchronous calls
	args := &Args2{7, 8}
	reply := new(Reply2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := client.Call(ctx, "Airth2.Add", args, reply); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBirpcInContextReflect(b *testing.B) {
	newServer := rpcc.NewBirpcServer()
	newServer.Register(new(rpcc.Airth2))
	c1, c2 := net.Pipe()
	go newServer.ServeConn(c1)

	client := rpcc.NewBirpcClient(c2)
	defer client.Close()
	client.Register(new(rpcc.Airth2))
	ctx := context.Background()

	// Synchronous calls
	args := &rpcc.Args2{7, 8}
	reply := new(rpcc.Reply2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := client.Call(ctx, "Airth2.Add", args, reply); err != nil {
			b.Fatal(err)
		}
	}
}
