// server.go
package main

import (
	"errors"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Arith int

type Args struct{ A, B int }
type Quotient struct{ Quo, Rem int }

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("Divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	// 实例化Arith对象并注册到rpc服务
	rpc.Register(new(Arith))
	listen, _ := net.Listen("tcp", ":1234")
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}
