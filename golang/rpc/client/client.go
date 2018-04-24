package main

import (
	"net"
	"net/rpc/jsonrpc"
	"feilin.com/gocourse/golang/rpc/server"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div", server.Args{10, 3}, &result)
	if err != nil {
		fmt.Printf("error: %v", err)
	} else {
		fmt.Println(result)
	}

	err = client.Call("DemoService.Div", server.Args{10, 0}, &result)
	if err != nil {
		fmt.Printf("error: %v", err)
	} else {
		fmt.Println(result)
	}
}
