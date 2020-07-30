package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	b := []byte("client starts")

	fmt.Printf("%s\n", b)

	p, _ := json.Marshal(map[string]interface{}{})
	fmt.Printf("%s\n", p)

	cli := &http.Client{Timeout: time.Second}

	for now := range time.Tick(100 * time.Millisecond) {
		resp, err := cli.Get("http://127.0.0.1:8080/get")
		if err != nil {
			panic(err)
		}

		if resp.StatusCode != http.StatusOK {
			panic("not ok")
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		// resp.Body.Close()

		fmt.Printf("%v: get %v bytes \n", now, len(b))
	}
}
