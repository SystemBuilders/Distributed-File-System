package lockserver

import "fmt"

// Server implements a lockserver
type Server struct {
	// rpc stuff
	// lock maintaining stuff
}

// Service is the service exposed by this server to the clients
type Service int

// HealthCheck acts like a ping
func (s *Service) HealthCheck(ip string, counter *int) error {
	fmt.Println("HealthCheck ping")
	return nil
}

// Acquire function lets a client acquire a lock on an object.
func (s *Service) Acquire(ip int, counter *float32) error {
	fmt.Println("Implement me")
	return nil
}

// Release lets a client to release a lock on an object.
func (s *Service) Release(ip int, counter *float32) error {
	fmt.Println("Implement me too")
	return nil
}
