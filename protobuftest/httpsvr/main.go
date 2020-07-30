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
	"syscall"
	"time"

	pb "github.com/liangzuobin/sometest/protobuftest/demo"
	"google.golang.org/grpc"
)

var cli pb.StorageClient

func init() {
	cc, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	cli = pb.NewStorageClient(cc)
}

func main() {

	// go loopPut()

	go func() {
		http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
			resp, err := cli.Get(r.Context(), &pb.JustKey{Key: "foo"})
			if err != nil {
				log.Println(err)
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
	time.Sleep(time.Second)
}

func loopPut() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan struct{}, 1000)

	for {
		ch <- struct{}{}

		go func() {
			defer func() { <-ch }()

			if err := putFile(ctx); err != nil {
				log.Println(err)
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

	f, err := os.Open(p)
	if err != nil {
		return err
	}

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
