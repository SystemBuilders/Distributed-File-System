package lockservice

import (
	"log"
	"sync"

	"github.com/rs/zerolog"
)

// SafeLockMap is the lockserver's data structure
type SafeLockMap struct {
	LockMap map[string]bool
	Mutex   sync.Mutex
}

var _ LockService = (*simpleLockService)(nil)

type simpleLockService struct {
	log     zerolog.Logger
	lockMap SafeLockMap
}

var _ Descriptors = (*simpleDescriptor)(nil)

type simpleDescriptor struct {
}

func (sd *simpleDescriptor) ID() string {
	return ""
}

// Acquire function lets a client acquire a lock on an object.
func (ls *simpleLockService) Acquire(sd Descriptors) error {
	ls.lockMap.Mutex.Lock()
	if ls.lockMap.LockMap[sd.ID()] {
		ls.lockMap.Mutex.Unlock()
		return ErrFileAcquired
	}
	ls.lockMap.LockMap[sd.ID()] = true
	ls.lockMap.Mutex.Unlock()
	ls.
		log.
		Debug().
		Str("descriptor", sd.ID()).
		Msg("locked")
	return nil
}

// Release lets a client to release a lock on an object.
func (ls *simpleLockService) Release(sd Descriptors) error {
	ls.lockMap.Mutex.Lock()
	if ls.lockMap.LockMap[sd.ID()] {
		delete(ls.lockMap.LockMap, sd.ID())
		log.Printf("File: %v's lock released\n", sd.ID())
		ls.lockMap.Mutex.Unlock()
		return nil
	}
	ls.lockMap.Mutex.Unlock()
	return ErrCantReleaseFile
}

// CheckAcquire returns nil if the file is acquired
func (ls *simpleLockService) CheckAcquired(sd Descriptors) bool {
	ls.lockMap.Mutex.Lock()
	if ls.lockMap.LockMap[sd.ID()] {
		ls.lockMap.Mutex.Unlock()
		return true
	}
	ls.lockMap.Mutex.Unlock()
	return false
}

// CheckRelease returns nil if the file is released
func (ls *simpleLockService) CheckReleased(sd Descriptors) bool {
	ls.lockMap.Mutex.Lock()
	if ls.lockMap.LockMap[sd.ID()] {
		ls.lockMap.Mutex.Unlock()
		return false
	}
	ls.lockMap.Mutex.Unlock()
	return true
}
