package middleware

import "sync"

//stores requests counts fo each ip address
//and protects the shared map using a mutex.//
type RateLimiter struct {
	requests map[string]int
	mu       sync.Mutex
	//mutex prevent race conditions and race condition exploits//
}

//limiter is one shared rate limiter instance used.
//by the entire server to track incoming requests.//
//using a package level variable architecture//
var limiter = &RateLimiter{
	//make to prepare map /
	//allocate memory and prepare this map for use//
	requests: make(map[string]int),
}

//when app starts,start the
//  rate-limit reset system automatically//
//init here is a go reserved word to say run before main() runs//
func init() {

}
