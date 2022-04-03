// client.go
package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// stub
type Args struct{ A, B int }
type Quotient struct{ Quo, Rem int }

func main() {
	client, err := jsonrpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	// Synchronous call
	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

	// Asynchronous call
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done // will be equal to divCall
	if replyCall.Error != nil {
		log.Fatal("arith error:", replyCall.Error)
	}
	fmt.Printf("Arith: %d / %d = %d...%d", args.A, args.B, quotient.Quo, quotient.Rem)
	// check errors, print, etc.
}
