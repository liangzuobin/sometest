package main

import (
	"crypto/rand"
	"log"
)

func main() {
	var count0, count1 int

	for i := 0; i < 1000; i++ {
		switch n := foo(); n {
		case 0:
			count0++
		case 1:
			count1++
		default:
			log.Printf("sum got %v", n)
		}
	}

	log.Printf("count 0: %v, count 1: %v", count0, count1)
}

func foo() int {
	buf := make([]byte, 10)
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}

	var sum int
	for _, b := range buf {
		sum += int(b)
	}

	return sum & 1
}
