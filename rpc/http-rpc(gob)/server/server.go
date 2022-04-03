/*
 * @Author: FANG
 * @Date: 2022-04-03 18:42:59
 */
package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
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
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Println("error: ", err)
	}
	http.Serve(listen, nil)
}
