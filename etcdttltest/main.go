package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func newClient() (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 3 * time.Second,
	})
}

func main() {
	foo := func(words ...string) {
		for _, w := range words {
			log.Println(w)
		}
	}

	foo()

	cli, err := newClient()
	if err != nil {
		log.Printf("new client failed: %v", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for e := range cli.Watch(ctx, "/foo", clientv3.WithPrefix()) {
			log.Printf("watch: %#v", e)
		}
	}()

	go register(ctx)
	watchSignals()
	cancel()
}

func register(ctx context.Context) {
	log.Println("register run")

	cli, err := newClient()
	if err != nil {
		log.Printf("new client failed: %v", err)
		return
	}

	resp, err := cli.Lease.Grant(ctx, 2)
	if err != nil {
		log.Printf("lease grant failed: %v", err)
		return
	}

	id := resp.ID

	if _, err := cli.Lease.KeepAlive(ctx, id); err != nil {
		log.Printf("lease keep alive failed: %v", err)
		return
	}

	time.Sleep(5 * time.Second)

	// sess, err := concurrency.NewSession(cli, concurrency.WithTTL(10))
	// if err != nil {
	// 	log.Printf("new session failed: %v", err)
	// }

	log.Println("put /foo/bar")
	// if _, err := cli.Put(ctx, "/foo/bar", "bar", clientv3.WithLease(sess.Lease())); err != nil {
	if _, err := cli.Put(ctx, "/foo/bar", "bar", clientv3.WithLease(id)); err != nil {
		log.Printf("put bar failed: %v", err)
		// sess.Close()

		return
	}

	log.Println("put /foo/baz")
	if _, err := cli.Put(ctx, "/foo/baz", "baz"); err != nil {
		log.Printf("put baz failed: %v", err)
		return
	}

	time.Sleep(5 * time.Second)

	log.Println("lease revoked")
	if _, err := cli.Lease.Revoke(ctx, id); err != nil {
		log.Printf("lease revoke failed: %v", err)
		// sess.Close()

		return
	}

	// if _, err := cli.Delete(ctx, "/foo/bar"); err != nil {
	// 	log.Printf("del failed: %v", err)
	// }

	// sess.Close()

	// cli.Close()
	// log.Println("client closed")

	// time.Sleep(3 * time.Second)

	// sess.Close()
	// log.Println("session closed")

	time.Sleep(3 * time.Second)

	rev, err := cli.Get(ctx, "/foo", clientv3.WithPrefix())
	if err != nil {
		log.Printf("get /foo failed: %v", err)
		return
	}

	log.Printf("rev: %+v", rev)

	cli.Close()
	log.Println("client closed")

	time.Sleep(time.Minute)
}

func watchSignals() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("syscall: %v", <-ch)
}
