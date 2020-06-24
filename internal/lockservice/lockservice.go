package lockservice

import (
	"errors"
	"log"
	"sync"
)

// SafeLockMap is the lockserver's data structure
type SafeLockMap struct {
	LockMap map[string]bool
	Mutex   sync.Mutex
}

// Service is the service exposed by this server to the clients.
type Service int

var lockMap = newSafelockMap()

func newSafelockMap() (lockMap SafeLockMap) {
	lockMap = SafeLockMap{}
	lockMap.LockMap = make(map[string]bool)
	return
}

// CheckAcquire returns nil if the file is acquired
func CheckAcquire(fileID string) error {
	lockMap.Mutex.Lock()
	if lockMap.LockMap[fileID] {
		lockMap.Mutex.Unlock()
		return nil
	}
	lockMap.Mutex.Unlock()
	return errors.New("Lock hasn't been acquired")
}

// Acquire function lets a client acquire a lock on an object.
func Acquire(fileID string) error {
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

// CheckRelease returns nil if the file is released
func CheckRelease(fileID string) error {
	lockMap.Mutex.Lock()
	if lockMap.LockMap[fileID] {
		lockMap.Mutex.Unlock()
		return nil
	}
	lockMap.Mutex.Unlock()
	return errors.New("Lock wasn't acquired on file")
}

// Release lets a client to release a lock on an object.
func Release(fileID string) error {
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
