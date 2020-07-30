package main

import (
	"context"
	"crypto/rand"
	"io"
	"log"
	"net/url"
	"time"
)

type foo interface {
	Hello(string)
}

type bar struct{}

func (b *bar) Hello(string) {
	log.Println("bar")
}

var _ foo = &bar{}

type baz struct {
	*bar
}

var _ foo = &baz{}

type channer struct {
	ch chan struct{}
}

func main() {

	log.Printf("%v", []byte("1"))
	log.Printf("%v", randomBytesMod(6, 10))

	{
		ctx := context.WithValue(context.TODO(), "foo", "foo")
		ctx = context.WithValue(ctx, "bar", 1)
		ctx = context.WithValue(ctx, "baz", struct{}{})
		log.Printf("foo: %v", value(ctx, "foo"))
		log.Printf("bar: %v", value(ctx, "bar"))
		log.Printf("baz: %v", value(ctx, "baz"))
		log.Printf("notexists: %v", value(ctx, "notexists"))
	}

	var ch channer
	ch.ch = make(chan struct{})

	go func() {
		select {
		case <-ch.ch:
			log.Println("receive from nil chan")
		}
	}()

	close(ch.ch)
	// log.Printf("closed chan is nil: %v", ch.ch == nil) // false
	ch.ch = nil
	time.Sleep(time.Millisecond)

	if s := toString(nil); s == "" {
		log.Println("nil to string")
		return
	}

	b := &bar{}
	b.Hello("string")

	testURLValues()

	testCap()

	var n *bar

	testNil(n)

	var s []*bar

	log.Printf("s is nil: %v", s == nil)

	for _, v := range s {
		log.Printf("ghost in s %v", v)
	}
}

func testNil(i interface{}) {
	f, ok := i.(*bar)
	if !ok {
		log.Println("nil cast to ptr not ok")
		return
	}

	log.Printf("nil cast to ptr %v", f)
}

func testURLValues() {
	v := url.Values{}

	v.Add("foo", "hello")
	v.Add("bar", "world")
	v.Add("baz", "asshole")

	s, err := url.QueryUnescape(v.Encode())
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("unescaped url = %s", s)

	log.Printf("escaped url = %s", url.QueryEscape(v.Encode()))
}

func testCap() {
	b := make([]byte, 0, 10)

	log.Printf("cap %v, len %v", cap(b), len(b))

	for i := 0; i < 9; i++ {
		b = append(b, 'q')
	}
	log.Printf("cap %v, len %v", cap(b), len(b))

	log.Printf("%v", b[0:9])
}

func toString(i interface{}) string {
	if s, ok := i.(string); ok {
		return s
	}
	return ""
}

func value(ctx context.Context, key string) string {
	if s, ok := ctx.Value(key).(string); ok {
		return s
	}
	return ""
}

func randomBytesMod(length int, mod byte) (b []byte) {
	if length == 0 {
		return nil
	}
	if mod == 0 {
		panic("captcha: bad mod argument for randomBytesMod")
	}
	maxrb := 255 - byte(256%int(mod))
	b = make([]byte, length)
	i := 0
	for {
		r := randomBytes(length + (length / 4))
		for _, c := range r {
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = c % mod
			i++
			if i == length {
				return
			}
		}
	}

}

func randomBytes(length int) (b []byte) {
	b = make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic("captcha: error reading random source: " + err.Error())
	}
	return
}
