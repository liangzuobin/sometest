package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
)

var m = make(map[string]int, 16)

func main() {

	for _, n := range []string{
		"/Users/liangzuobin/Downloads/storage-provider.error.log",
		"/Users/liangzuobin/Downloads/storage-provider.error.log.20200709000006",
		"/Users/liangzuobin/Downloads/storage-provider.error.log.20200709123754",
		"/Users/liangzuobin/Downloads/storage-provider.error.log.20200709123804",
	} {
		readfile(n)
	}

	log.Printf("total: %v", len(m))

	for k, v := range m {
		if v > 1000 {
			log.Printf("%v: %v", k, v)
		}
	}
}

var reg = regexp.MustCompile(`key="(.*)"`)

func readfile(name string) {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(f)

	for {
		switch ln, _, err := r.ReadLine(); err {
		case nil:
			if s := reg.FindAllSubmatch(ln, 1); len(s) > 0 {
				// log.Println(string(s[0][1]))
				m[string(s[0][1])]++
			}
		case io.EOF:
			return
		default:
			log.Panic(err)
		}
	}
}
