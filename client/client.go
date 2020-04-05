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

	var x int
	var y float32
	x = 2
	y = 3.4

	err = client.Call("Service.HealthCheck", "10.100.20.23", &x)
	if err != nil {
		fmt.Println(err)
	}
	err = client.Call("Service.Acquire", 1, &y)
	if err != nil {
		fmt.Println(err)
	}
}
