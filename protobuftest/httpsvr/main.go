package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	pb "github.com/liangzuobin/sometest/protobuftest/demo"
	"google.golang.org/grpc"
)

type clientFunc func(ctx context.Context, unary bool) pb.StorageClient

var cli clientFunc

func mustNewClient(ctx context.Context) pb.StorageClient {
	cc, err := grpc.DialContext(ctx, "10.13.254.37:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return pb.NewStorageClient(cc)
}

func singleClient(ctx context.Context) clientFunc {
	cli := mustNewClient(ctx)

	return func(context.Context, bool) pb.StorageClient {
		return cli
	}
}

func separateClient(ctx context.Context) clientFunc {
	unr, str := mustNewClient(ctx), mustNewClient(ctx)

	return func(ctx context.Context, unary bool) pb.StorageClient {
		if unary {
			return unr
		}
		return str
	}
}

func pooledClient(ctx context.Context) clientFunc {
	var next int64
	pool := []pb.StorageClient{mustNewClient(ctx), mustNewClient(ctx)}
	var mu sync.Mutex

	return func(context.Context, bool) pb.StorageClient {

		// nn := atomic.AddInt64(&next, 1)

		mu.Lock()
		defer mu.Unlock()
		next++
		nn := next

		return pool[nn&1] // FIXME 因为只有两个 client 才这么处理
	}
}

// test it with `$ ./go-wrk -c=100 -t=5 -k=true -n=1000 http://127.0.0.1:8090/get`

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// 每次调用 client 返回相同的连接
	// cli = singleClient(ctx)

	// 所有 unary 请求共享一个 client；stream 请求共用一个 client
	cli = separateClient(ctx)

	// 每次请求，创建新的链接 // 有 bug 创建完了没有关闭掉
	// cli = func(ctx context.Context, _ bool) pb.StorageClient {
	// 	return mustNewClient(ctx)
	// }

	// 从一个只有两个链接的链接池里拿
	// cli = pooledClient(ctx)

	var once sync.Once

	go func() {

		http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
			once.Do(func() {
				go loopPut(ctx, 200) // 启用干扰用的 stream
			})

			resp, err := cli(ctx, true).Get(r.Context(), &pb.JustKey{Key: "foo"})
			if err != nil {
				log.Printf("/get failed: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Fprintln(w, resp.Bytes)
		})

		_ = http.ListenAndServe(":8090", nil)
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	cancel()
	time.Sleep(time.Second)
}

func loopPut(ctx context.Context, n int) {
	errors := make(chan error, n)

	go func() {
		m := make(map[string]int, 32)

		defer func() {
			log.Printf("errors count %v", len(m))

			for k, v := range m {
				log.Printf("err: %v \ncount %v", k, v)
			}
		}()

		for err := range errors {
			select {
			case <-ctx.Done():
				return
			default:
				m[err.Error()]++
			}
		}
	}()

	ch := make(chan struct{}, n)

	for {
		select {
		case <-ctx.Done():
			return
		case ch <- struct{}{}:
		}

		go func() {
			defer func() { <-ch }()

			if err := put(ctx); err != nil {
				errors <- err
			}
		}()
	}
}

func put(ctx context.Context) error {
	str, err := cli(ctx, false).Put(ctx)
	if err != nil {
		return err
	}

	if err := str.Send(&pb.StreamReq{Key: "hello"}); err != nil {
		return err
	}

	// 每次 32k 数据，发 10 次
	p := make([]byte, 32*1024)

	for i := 0; i < 10; i++ {
		n, err := rand.Reader.Read(p)
		if err != nil {
			_ = str.CloseSend()
			return err
		}

		if err := str.Send(&pb.StreamReq{Value: p[:n]}); err != nil {
			return err
		}
	}

	if _, err := str.CloseAndRecv(); err != nil {
		return err
	}

	return nil
}
