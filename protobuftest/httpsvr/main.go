package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	pb "github.com/liangzuobin/sometest/protobuftest/demo"
	"google.golang.org/grpc"
)

var cli pb.StorageClient

// test it with `$ ./go-wrk -c=100 -t=5 -k=true -n=1000 http://127.0.0.1:8090/get`

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cc, err := grpc.DialContext(ctx, "10.13.254.37:9080", grpc.WithInsecure())
	if err != nil {
		cancel()
		log.Fatal(err)
	}

	cli = pb.NewStorageClient(cc)

	var once sync.Once

	go func() {
		cc, err := grpc.DialContext(ctx, "10.13.254.37:9080", grpc.WithInsecure())
		if err != nil {
			cancel()
			log.Fatal(err)
		}

		otherCli := pb.NewStorageClient(cc)

		http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
			once.Do(func() {
				go loopPut(ctx) // 启用干扰用的 stream
			})

			// resp, err := cli.Get(r.Context(), &pb.JustKey{Key: "foo"})
			resp, err := otherCli.Get(r.Context(), &pb.JustKey{Key: "foo"})
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
	time.Sleep(3 * time.Second)
}

func loopPut(ctx context.Context) {
	n := 300

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

	ratio := make(chan struct{}, n)

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		ratio <- struct{}{}

		go func() {
			defer func() { <-ratio }()

			if err := putFile(ctx); err != nil {
				errors <- err
			}
		}()
	}
}

func putFile(ctx context.Context) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	p := filepath.Join(wd, "main.go")

	str, err := cli.Put(ctx)
	if err != nil {
		return err
	}

	if err := str.Send(&pb.StreamReq{Key: p}); err != nil {
		return err
	}

	buf := make([]byte, 1024)

	send := func(n int) error {
		return str.Send(&pb.StreamReq{Value: buf[:n]})
	}

	f, err := os.Open(p)
	if err != nil {
		return err
	}
	defer f.Close()

	for {
		switch n, err := f.Read(buf); err {
		case nil:
			if err := send(n); err != nil {
				return err
			}
		case io.EOF:
			if n > 0 {
				if err := send(n); err != nil {
					return err
				}
			}

			if _, err := str.CloseAndRecv(); err != nil {
				return err
			}

			return nil
		default:
			return err
		}
	}
}
