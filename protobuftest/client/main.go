package main

import (
	"context"
	"crypto/rand"
	"io"
	"log"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	pb "github.com/liangzuobin/sometest/protobuftest/demo"
	"google.golang.org/grpc"
)

var succeed, failed int32

var m sync.Map

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "127.0.0.1:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	cli := pb.NewStorageClient(conn)

	// for j := 0; j < 10; j++ {
	// 	var wg sync.WaitGroup

	// 	for i := 0; i < 1000; i++ {
	// 		wg.Add(1)

	// 		go func(wg *sync.WaitGroup) {
	// 			get(cli, i)
	// 			wg.Done()
	// 		}(&wg)
	// 	}

	// 	wg.Wait()

	// 	time.Sleep(3 * time.Second)
	// 	log.Printf("batch: %v, succeed: %v, failed: %v", j, succeed, failed)
	// }

	// m.Range(func(k, v interface{}) bool {
	// 	log.Printf("%v", k)
	// 	return false
	// })

	// 	tk := time.Tick(time.Second)

	// 	var wg sync.WaitGroup
	// 	for j := 0; ; j++ {
	// 		select {
	// 		case <-tk:
	// 			goto wati
	// 		default:
	// 		}

	// 		wg.Add(1)
	// 		go func(j int) {
	// 			put(cli, j)
	// 			wg.Done()
	// 		}(j)
	// 	}
	// wati:
	// 	wg.Wait()

	if err := put2(cli); err != nil {
		panic(err)
	}

}

func get(cli pb.StorageClient, i int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if _, err := cli.Get(ctx, &pb.JustKey{Key: strconv.Itoa(i)}); err != nil {
		atomic.AddInt32(&failed, 1)
		m.Store(err, "")

		return
	}

	atomic.AddInt32(&succeed, 1)
}

func put(cli pb.StorageClient, i int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := cli.Put(ctx)
	if err != nil {
		panic(err)
	}

	if err := stream.Send(&pb.StreamReq{Key: strconv.Itoa(i)}); err != nil {
		panic(err)
	}

	b := make([]byte, 10)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	if err := stream.Send(&pb.StreamReq{Value: b}); err != nil {
		panic(err)
	}

	switch _, err := stream.CloseAndRecv(); err {
	case nil:
	case io.EOF:
		log.Println("put receive io.EOF")
	default:
		panic(err)
	}
}

func put2(cli pb.StorageClient) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := cli.Put2(ctx)
	if err != nil {
		return err
	}

	defer func() {
		log.Printf("grpc got err: %v", err)
		_, err = stream.CloseAndRecv()
		log.Printf("real err: %v", err)
	}()

	for i := 0; i < 1000; i++ {
		b := make([]byte, 8)
		if _, err := rand.Read(b); err != nil {
			log.Printf("streaming failed: %v", err)
			return err
		}

		if err := stream.Send(&pb.StreamReq{Key: strconv.Itoa(i), Value: b}); err != nil {
			log.Printf("streaming failed: %v", err)
			return err
		}
	}

	// switch _, err := stream.CloseAndRecv(); err {
	// case nil, io.EOF:
	// 	return nil
	// default:
	// 	return err
	// }
	return
}
