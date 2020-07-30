package main

import (
	"math/rand"
	"testing"
)

func TestRandomBytes(t *testing.T) {
	for i := 0; i < 999; i++ {
		v := randomBytesMod(6, 10)
		t.Logf("%v: %06d", v, rand.Intn(999999))
	}
}
