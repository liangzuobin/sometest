package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/liangzuobin/sometest/protobuftest/demo"
	"google.golang.org/grpc"
)

func main() {
	foo := "helloworld where are you from i am fine thank you and you?"
	if len(foo) != len([]byte(foo)) {
		log.Fatal("NOT EQUALS")
	}

	lis, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatal(err)
	}

	// svr := grpc.NewServer(grpc.MaxConcurrentStreams(10))
	svr := grpc.NewServer()
	pb.RegisterStorageServer(svr, &server{})

	go func() {
		if err := svr.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		http.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(200 * time.Millisecond)
			fmt.Fprint(w, "ok")
		})
		http.ListenAndServe(":8080", nil)
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	svr.GracefulStop()
	time.Sleep(time.Second)
}

const simulateHandlingCost = 200 * time.Millisecond

type server struct{}

func (s *server) Get(_ context.Context, _ *pb.JustKey) (*pb.JustBytes, error) {
	time.Sleep(simulateHandlingCost)
	return &pb.JustBytes{Bytes: []byte("foo")}, nil
}

var client = http.Client{}

type putReader struct {
	ctx    context.Context
	stream pb.Storage_PutServer
	buf    []byte
	index  int
}

func (r *putReader) Read(p []byte) (n int, err error) {
	select {
	case <-r.ctx.Done():
		return 0, r.ctx.Err()
	default:
	}

	if r.index == len(r.buf) {
		msg, err := r.stream.Recv()
		if err != nil {
			return 0, err
		}
		r.buf = msg.Value
		r.index = 0
	}

	for n < cap(p) && r.index < len(r.buf) {
		p[n] = r.buf[r.index]
		n++
		r.index++
	}

	return n, nil
}

func (s *server) Put(stream pb.Storage_PutServer) error {
	ctx := stream.Context()

	md, err := stream.Recv()
	if err != nil {
		log.Printf("recv md got: %v, %v", md, err)
		return err
	}

	// log.Printf("receive: %v", md.Key)

	// 这里返回 error 不会导致客户端 EOF
	// if i, err := strconv.Atoi(md.Key); err == nil && i%10 == 0 {
	// 	// stream.SendAndClose(nil)
	// 	return fmt.Errorf("screw you: %v", i)
	// }

	rd := &putReader{
		ctx:    ctx,
		stream: stream,
	}

	if _, err := io.Copy(ioutil.Discard, rd); err != nil {
		return err
	}

	// req, err := http.NewRequest("PUT", "http://127.0.0.1:8080/put", rd)
	// if err != nil {
	// 	log.Printf("new req got: %v", err)
	// 	return err
	// }

	// resp, err := client.Do(req.WithContext(ctx))
	// if err != nil {
	// 	log.Printf("put got: %v", err)
	// 	return err
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	return fmt.Errorf("not ok: %v", resp.Status)
	// }

	return stream.SendAndClose(&pb.JustBytes{}) // 这样 client 在执行 CloseAndRecv() 的时候，就不会报 io.EOF 了
}

func (s *server) Put2(stream pb.Storage_Put2Server) error {
	ctx := stream.Context()

	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if i == 5 {
			stream.SendAndClose(&pb.JustBytes{})
			// return nil
			return fmt.Errorf("screw you: %v", i)
			// break
		}

		msg, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("recevie: %v", msg.Key)
	}

	return stream.SendAndClose(&pb.JustBytes{}) // 这样 client 在执行 CloseAndRecv() 的时候，就不会报 io.EOF 了
}
