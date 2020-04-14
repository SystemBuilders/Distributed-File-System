package lockserver

import (
	"fmt"
	"log"
	"net/rpc"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSingleLockAndRelease(t *testing.T) {
	go StartServer()
	timer1 := time.NewTimer(5 * time.Millisecond)
	<-timer1.C
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
}
