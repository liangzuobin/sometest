package main

import (
	"errors"
	"log"
	"time"
)

func main() {
	ch := make(chan error)

	go func() {
		time.Sleep(time.Second)
		ch <- errors.New("hello")
		time.Sleep(time.Second)
		close(ch)
	}()

	log.Println("waiting for ch returns err")
	log.Println(<-ch)

	log.Println("waiting for ch to be closed")
	v, ok := <-ch
	log.Println(v, ok)
	time.Sleep(time.Second)
}
