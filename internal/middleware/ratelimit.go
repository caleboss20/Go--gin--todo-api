package middleware

import (
	"sync"
	"time"
)

//this is a fixed Window (Rate limiting) Counter Algorith//

// stores requests counts fo each ip address
// and protects the shared map using a mutex.//
type RateLimiter struct {
	requests map[string]int
	mu       sync.Mutex
	//mutex prevent race conditions and race condition exploits//
	//eg====>prevent TOCTOU(time of check to time of use)
}

// limiter is one shared rate limiter instance used.
// by the entire server to track incoming requests.//
// using a package level variable architecture//
var limiter = &RateLimiter{
	//make to prepare map /
	//allocate memory and prepare this map for use//
	requests: make(map[string]int),
}

// when app starts,start the
//
//	rate-limit reset system automatically//
//
// init here is a go reserved word to say run before main() runs//
// resets request counts every minute.//
func init() {
	//create anonymous func,immediately run it//
	//Run it as a goroutine in background//

	go func() {
		for {
			time.Sleep(time.Minute)
			limiter.mu.Lock()
			//reset all request counts//
			limiter.requests = make(map[string]int)
			limiter.mu.Unlock()
		}
	}()
}
