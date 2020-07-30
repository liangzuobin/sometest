package main

import (
	"crypto/rand"
	"log"
	mr "math/rand"
	"sync"
)

type itf interface {
	foo()
}

type cache struct {
	// mu sync.Mutex
	mu sync.RWMutex
	m  map[string]bool
}

func main() {
	// foo()
	// bar()

	var i itf

	if i == nil {
		return
	}

	baz()
}

func foo() {
	c := cache{m: make(map[string]bool, 1024)}

	var wg sync.WaitGroup

	for i := 0; i < 1024; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.mu.Lock()
			c.m[key()] = true
			c.mu.Unlock()
		}()

		wg.Add(1)

		go func() {
			defer wg.Done()
			c.mu.RLock()
			log.Println(c.m[key()])
			c.mu.RUnlock()
		}()

	}

	wg.Wait()
}

func bar() {
	c := &cache{m: make(map[string]bool, 1024)}

	var wg sync.WaitGroup

	for i := 0; i < 1024; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.mu.Lock()
			c.m[key2()] = true
			c.mu.Unlock()
		}()

		wg.Add(1)

		go func() {
			defer wg.Done()
			c.mu.RLock()
			if !c.m[key2()] {
				log.Println(false)
			}
			c.mu.RUnlock()
		}()

	}

	wg.Wait()
}

func baz() {
	c := &cache{m: make(map[string]bool, 1024)}

	c.mu.RLock()
	log.Println("got r lock.")

	c.mu.Lock()
	log.Println("got w lock.")

	c.mu.Unlock()
	log.Println("release w lock.")

	c.mu.RUnlock()
	log.Println("done.")
}

func key() string {
	b := make([]byte, 3)
	if _, err := rand.Read(b); err != nil {
		log.Fatalf("generate random key failed: %v", err)
	}

	return string(b)
}

func key2() string {
	if mr.Int()&1 == 0 {
		return "hello"
	}
	return "world"
}
