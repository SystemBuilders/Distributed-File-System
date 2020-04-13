package lockserver

import (
	"fmt"
	"log"
	"net/rpc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoubleLockAndRelease(t *testing.T) {
	t.Logf("acquire a release a acquire a release a")
	// go StartServer()
	assert := assert.New(t)
	client, err := rpc.DialHTTP("tcp", "localhost:55550")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	var y float32
	y = 3.4

	err = client.Call("Service.Acquire", "1", &y)

	if err != nil {
		assert.Fail(fmt.Sprintf("Failed to acquire lock: %v", err))
	}
	err = client.Call("Service.Release", "1", &y)
	if err != nil {
		assert.Fail(fmt.Sprintf("Failed to release lock: %v", err))
	}

	err = client.Call("Service.Acquire", "1", &y)

	if err != nil {
		assert.Fail(fmt.Sprintf("Failed to acquire lock: %v", err))
	}
	err = client.Call("Service.Release", "1", &y)
	if err != nil {
		assert.Fail(fmt.Sprintf("Failed to release lock: %v", err))
	}

	t.Logf("PASSED")

}

func TestMultipleLockAndRelease(t *testing.T) {
	t.Logf("acquire a acquire b release b release a\n")
	assert := assert.New(t)
	client, err := rpc.DialHTTP("tcp", "localhost:55550")

	var y float32
	y = 3.4
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	err = client.Call("Service.Acquire", "1", &y)

	if err != nil {
		assert.Fail(fmt.Sprintf("Failed to acquire lock: %v", err))
	}
	err = client.Call("Service.Acquire", "2", &y)
	if err != nil {
		assert.Fail(fmt.Sprintf("Failed to release lock: %v", err))
	}

	err = client.Call("Service.Release", "2", &y)
	if err != nil {
		assert.Fail(fmt.Sprintf("Failed to release lock: %v", err))
	}

	err = client.Call("Service.Release", "1", &y)
	if err != nil {
		assert.Fail(fmt.Sprintf("Failed to release lock: %v", err))
	}

}
