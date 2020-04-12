package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func errorCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func test1(client *Client) {
	var y float32
	y = 3.4
	fmt.Println("running Test1: ")
	fmt.Println("acquire	 a release a acquire a release a\n")
	err := client.Call("Service.Acquire", "a", &y)
	errorCheck(err)
	err = client.Call("Service.Release", "a", &y)
	errorCheck(err)
	err = client.Call("Service.Acquire", "a", &y)
	errorCheck(err)
	err = client.Call("Service.Release", "a", &y)
	errorCheck(err)

	fmt.Println("Test1 PASSED")

}

func main() {

	// Connect to lockserver
	client, err := rpc.DialHTTP("tcp", "localhost:55550")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	test1(client)

}
