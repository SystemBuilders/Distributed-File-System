package lockserver

import (
	"errors"
	"fmt"
	"log"
	"sync"
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
