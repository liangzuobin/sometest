package grpcpool

import (
	"testing"
)

func TestPoolAdd(t *testing.T) {
	p := &listConnsPool{}
	addr := "foo"
	cnn1 := &conn{next: emptyConn}
	cnn2 := &conn{next: emptyConn}
	cnn3 := &conn{next: emptyConn}

	p.add(addr, cnn1)
	p.add(addr, cnn2)
	p.add(addr, cnn3)

	v, ok := p.m.Load(addr)
	if !ok {
		t.Error("no conn exists")
		return
	}

	cnn := v.(*conn)

	if cnn.next == nil {
		t.Error("cnn should has next")
		return
	}

	if cnn.next != cnn2 {
		t.Error("cnn.next should be cnn1")
		return
	}

	if cnn.next.next == nil {
		t.Error("cnn.next.next should has next")
		return
	}

	if cnn.next.next != cnn3 {
		t.Error("cnn.next.next should has next")
		return
	}
}
