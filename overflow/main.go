package main

import (
	"log"
	"math"
)

func main() {
	var n uint8 = math.MaxUint8

	for i := 0; i < 10; i++ {
		log.Println(n)
		n++
	}
}
