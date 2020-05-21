package main

import (
	"log"
	"net"
	"time"
)

func main() {
	raw := ":0"

	ch := make(chan struct{})

	go func() {
		ln, err := net.Listen("tcp", raw)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("ln: %s", ln.Addr().String())

		<-ch
		ln.Close()
	}()

	go func() {
		ln2, err := net.Listen("tcp", raw)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("ln2: %s", ln2.Addr().String())

		<-ch
		ln2.Close()
	}()

	time.Sleep(time.Second)
	close(ch)
}
