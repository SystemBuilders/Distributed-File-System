package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:55550")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	var y float32
	y = 3.4

	err = client.Call("Service.Acquire", "1", &y)
	if err != nil {
		fmt.Println(err)
	}
	err = client.Call("Service.Release", "1", &y)
	if err != nil {
		fmt.Println(err)
	}
}
