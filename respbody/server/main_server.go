package main

import (
	"crypto/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		buf := make([]byte, 1024)
		if _, err := rand.Read(buf); err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(buf)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
