package main

import (
	"context"
	"log"
	"time"
)

type busy struct {

	// close 可以做到通知全部监听者的目的
	// closed chan struct{}

	ctx    context.Context
	cancel context.CancelFunc
}

// func (b *busy) loop1(ctx context.Context) {
// t := time.NewTicker(time.Second)

// for {
// 	select {
// 	case <-t.C:
// 		log.Println("loop1...")
// 	case <-b.closed:
// 		t.Stop()
// 		log.Println("loop1 closed")

// 		return
// 	}
// }

// }

// func (b *busy) loop2() {
// 	t := time.NewTicker(time.Second)

// 	for {
// 		select {
// 		case <-t.C:
// 			log.Println("loop2...")
// 		case <-b.closed:
// 			t.Stop()
// 			log.Println("loop2 closed")

// 			return
// 		}
// 	}
// }

// func (b *busy) close() {
// 	// b.closed <- struct{}{}
// 	close(b.closed)
// }

func (*busy) loop3() {
	for {
		log.Println("loop3...")
		time.Sleep(time.Second)
	}
}

func (b *busy) loop4() {
	for {
		select {
		case <-b.ctx.Done():
			log.Println("loop4 closed")

			return
		default:
		}

		log.Println("loop4...")
		time.Sleep(time.Second)
	}
}

func (b *busy) loop5() {
	t := time.NewTicker(time.Second)

	for {
		select {
		case <-t.C:
			log.Println("loop5...")
		case <-b.ctx.Done():
			t.Stop()
			log.Println("loop5 closed")

			return
		}
	}
}

func (b *busy) close() {
	log.Println("invoke close()")
	b.cancel()
}

func main() {
	// b := &busy{closed: make(chan struct{})}
	// go b.loop1()
	// go b.loop2()

	b := &busy{}

	b.ctx, b.cancel = context.WithCancel(context.Background())

	go b.loop3()
	go b.loop4()
	go b.loop5()

	const d = 5 * time.Second

	time.Sleep(d)

	b.close()

	time.Sleep(d)
	b.close() // 重复调用 cancel() 并没有什么关系
	time.Sleep(d)
}
