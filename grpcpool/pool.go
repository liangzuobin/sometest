package grpcpool

import (
	"context"
	"errors"
	"log"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

var ErrPoolClosed = errors.New("pool closed")

type Config struct{}

type Pool interface {
	Get(context.Context, string) (*grpc.ClientConn, error)
	Put(*grpc.ClientConn)
	Close() error
}

func NewPool(c Config) Pool {
	return nil
}

type conn struct {
	cli       *grpc.ClientConn
	ref       int32
	lastVisit atomic.Value
	next      *conn
}

var emptyConn = &conn{}

func (c *conn) hasNext() bool {
	return c.next != nil && c.next != emptyConn
}

func (c *conn) countGet() {
	atomic.AddInt32(&c.ref, 1)
	c.lastVisit.Store(time.Now())
}

func (c *conn) countPut() {
	atomic.AddInt32(&c.ref, -1)
	c.lastVisit.Store(time.Now())
}

func (c *conn) alive() bool {
	switch c.cli.GetState() {
	case connectivity.TransientFailure, connectivity.Shutdown:
		return false
	default:
		return true
	}
}

func (c *conn) idleSince(t time.Time) bool {
	return c.ref == 0 && c.lastVisit.Load().(time.Time).Before(t)
}

type listConnsPool struct {
	closed bool
	m      sync.Map
	ch     chan struct{} // 用来限制最多可以同时发起多少个 Dial 动作
}

func newListConnsPool(c Config) Pool {
	return &listConnsPool{ch: make(chan struct{}, 100)}
}

func (p *listConnsPool) Get(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	return nil, errors.New("fixme")
}

func (p *listConnsPool) get(ctx context.Context, addr string) (*conn, error) {
	if v, ok := p.m.Load(addr); ok {
		if cnn := v.(*conn); cnn.alive() {
			return cnn, nil
		}
	}

	p.ch <- struct{}{}
	defer func() { <-p.ch }()

	if v, ok := p.m.Load(addr); ok {
		if cnn := v.(*conn); cnn.alive() {
			return cnn, nil
		}
	}

	cli, err := grpc.DialContext(ctx, addr)
	if err != nil {
		return nil, err
	}

	p.add(addr, &conn{cli: cli, next: emptyConn})

	return p.get(ctx, addr)
}

func (p *listConnsPool) add(addr string, cnn *conn) {
	act, loaded := p.m.LoadOrStore(addr, cnn)
	if !loaded {
		if act.(*conn) == cnn {
			log.Println("hola")
		}
		return // store succeed
	}

	cp := unsafe.Pointer(cnn)

	for pred := act.(*conn); pred != nil; pred = pred.next {
		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&pred.next)),
			unsafe.Pointer(emptyConn),
			cp,
		) {
			log.Println("cas")
			return
		}
	}

	panic("impossible")
}

func (p *listConnsPool) Put(*grpc.ClientConn) {

}

func (p *listConnsPool) Close() error {
	if !p.closed {
		return nil
	}

	p.m.Range(func(k, v interface{}) bool {
		// FIXME
		return true
	})

	return nil
}
