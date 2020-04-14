package lockserver

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

// SafeLockMap is the lockserver's data structure
type SafeLockMap struct {
	LockMap map[string]bool
	Mutex   sync.Mutex
}

// Service is the service exposed by this server to the clients
type Service int

var lockMap = newSafelockMap()

func newSafelockMap() (lockMap SafeLockMap) {
	lockMap = SafeLockMap{}
	lockMap.LockMap = make(map[string]bool)
	return
}

// HealthCheck acts like a ping
func (s *Service) HealthCheck(ip string, counter *int) error {
	fmt.Println("HealthCheck ping")
	return nil
}

// checkAcquire returns true if the file is acquired
func (s *Service) checkAcquire(fileID string) bool {
	lockMap.Mutex.Lock()
	if lockMap.LockMap[fileID] {
		lockMap.Mutex.Unlock()
		return true
	}
	lockMap.Mutex.Unlock()
	return false
}

// Acquire function lets a client acquire a lock on an object.
func (s *Service) Acquire(fileID string, counter *float32) error {
	lockMap.Mutex.Lock()
	if lockMap.LockMap[fileID] {
		lockMap.Mutex.Unlock()
		return errors.New("Can't access file, locked by other user")
	}
	lockMap.LockMap[fileID] = true
	lockMap.Mutex.Unlock()
	log.Printf("File: %v locked\n", fileID)
	return nil
}

// checkRelease returns true if the file is released
func (s *Service) checkRelease(fileID string) bool {
	lockMap.Mutex.Lock()
	if lockMap.LockMap[fileID] {
		lockMap.Mutex.Unlock()
		return false
	}
	lockMap.Mutex.Unlock()
	return true
}

// Release lets a client to release a lock on an object.
func (s *Service) Release(fileID string, counter *float32) error {
	lockMap.Mutex.Lock()
	if lockMap.LockMap[fileID] {
		delete(lockMap.LockMap, fileID)
		log.Printf("File: %v's lock released\n", fileID)
		lockMap.Mutex.Unlock()
		return nil
	}
	lockMap.Mutex.Unlock()
	return errors.New("Can't release lock on file, wasn't locked before")
}

// StartServer starts the rpc server for lockserver
func StartServer() {
	service := new(Service)
	port := flag.Int("port", 55550, "service port")
	ip := flag.String("ip", "0.0.0.0", "service ip")

	rpcServer := rpc.NewServer()
	err := rpcServer.Register(service)
	if err != nil {
		log.Fatalf("Error in registering the RPC server: %v\n", err)
	}

	server := &http.Server{
		Addr:    *ip + ":" + strconv.Itoa(*port),
		Handler: rpcServer,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("listen error:", err)
		}
	}()

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	server.Shutdown(ctx)
	cancel()
}
