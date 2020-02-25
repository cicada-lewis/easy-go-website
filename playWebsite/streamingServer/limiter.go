package main

import "log"

type connLimiter struct {
	concurrentConn int
	bucket chan int
}

func NewConnLimiter(cc int) *connLimiter{
	return &connLimiter{
		concurrentConn: cc,
		bucket: make(chan int, cc),
	}
}

func (cl *connLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Print("Reached the rate limitation.")
		return false
	}

	cl.bucket <- 1
	return true
}

func (cl *connLimiter) ReleaseConn() {
	c := <- cl.bucket
	log.Printf("New Connection coming: %d", c)
}