package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

type foo struct {
	mu  sync.Mutex
	bar int64
	baz int64
}

func (f *foo) countByMu() {
	f.mu.Lock()
	f.bar++
	f.baz++
	f.mu.Unlock()
}

func (f *foo) countByAtomic() {
	atomic.AddInt64(&f.bar, 1)
	atomic.AddInt64(&f.baz, 1)
}

func BenchmarkCountByMu(b *testing.B) {
	f := new(foo)
	b.ReportAllocs()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			f.countByMu()
		}
	})
}

func BenchmarkCountByAtomic(b *testing.B) {
	f := new(foo)
	b.ReportAllocs()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			f.countByAtomic()
		}
	})
}
