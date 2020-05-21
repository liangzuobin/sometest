package main

import (
	"context"
	"log"
	"sync/atomic"

	"time"
)

var status uint32

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			log.Printf("status: %d", status)
		}
	}()

	time.Sleep(10 * time.Millisecond)
	log.Println("main change status to 1")
	status = 1

	time.Sleep(10 * time.Millisecond)
	log.Println("main change status to 2 atomically")
	atomic.StoreUint32(&status, 2)

	time.Sleep(10 * time.Millisecond)
	cancel()

	foo()
}

func foo(b ...int) {
	b = append(b, 1, 2, 3)
	log.Printf("%v", b)
}
