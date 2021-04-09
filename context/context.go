package context

import (
	"context"
	"time"
)

var (
	Canceled         = context.Canceled
	DeadlineExceeded = context.DeadlineExceeded
	background       = Context{Context: context.Background()}
	todo             = Context{Context: context.TODO()}
)

func Background() Context {
	return background
}

func TODO() Context {
	return todo
}

// ClientConnector is the connection used in RpcClient, as interface so we can combine the rpc.RpcClient with http one or websocket
type ClientConnector interface {
	Call(ctx Context, serviceMethod string, args, reply interface{}) error
}

type CancelFunc = context.CancelFunc

type Context struct {
	context.Context

	Client ClientConnector
}

func WithCancel(parent Context) (Context, CancelFunc) {
	var cancel CancelFunc
	parent.Context, cancel = context.WithCancel(parent.Context)
	return parent, cancel
}
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
	var cancel CancelFunc
	parent.Context, cancel = context.WithDeadline(parent.Context, d)
	return parent, cancel
}
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	var cancel CancelFunc
	parent.Context, cancel = context.WithTimeout(parent.Context, timeout)
	return parent, cancel
}

func WithValue(parent Context, key, val interface{}) Context {
	parent.Context = context.WithValue(parent.Context, key, val)
	return parent
}
