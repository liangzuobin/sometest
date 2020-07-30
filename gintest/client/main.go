package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
)

func main() {
	for {
		go user()

		runtime.Gosched()
	}
}

func user() {
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/user/1/2", nil)
	if err != nil {
		log.Printf("new req failed: %v", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("do req failed: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("%v", resp.Status)
}

func upload() {
	f, err := os.Open("/Users/liangzuobin/Downloads/avatar.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("put with content: %v", len(b))

	req, err := http.NewRequest(http.MethodPut, "http://127.0.0.1:8080/upload", bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Printf("%v", resp.Status)
}
