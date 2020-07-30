package main

import (
	"log"
	"sync"
	"sync/atomic"
	"unsafe"
)

type foo struct {
	s string
}

type bar struct {
	f *foo
}

func main() {
	casStruct()
	casSyncMap()
}

func casStruct() {
	b := &bar{}

	f := b.f
	log.Printf("b.f: %v", unsafe.Pointer(f))

	p := (*unsafe.Pointer)(unsafe.Pointer(&b.f))

	var (
		count int64
		wg    sync.WaitGroup
	)

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			n := unsafe.Pointer(new(foo))
			if ok := atomic.CompareAndSwapPointer(p, unsafe.Pointer(f), n); ok {
				log.Printf("cas: %v", ok)
				return
			}

			atomic.AddInt64(&count, 1)
		}()
	}
	wg.Wait()

	log.Printf("b: %+v, failed: %v times", b, count)
}

func casSyncMap() {
	var m sync.Map

	m.Store("key", new(foo))

	v, _ := m.Load("key")

	p := (*unsafe.Pointer)(unsafe.Pointer(&v))
	f := unsafe.Pointer(&v)

	var (
		wg    sync.WaitGroup
		count int64
	)

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			n := unsafe.Pointer(&foo{s: "foo"})
			if ok := atomic.CompareAndSwapPointer(p, f, n); ok {
				log.Printf("cas: %v", ok)
				return
			}

			atomic.AddInt64(&count, 1)
		}()
	}

	wg.Wait()

	v, _ = m.Load("key")

	log.Printf("v: %+v, failed: %v times", v, count)
}
