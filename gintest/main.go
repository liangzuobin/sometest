package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/user/:name/*action", func(c *gin.Context) {
		c.Set(key, "hello")
		c.Next()
	}, func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		how := c.Param("how")
		message := name + " is " + action + " how " + how
		go func() {
			// time.Sleep(10 * time.Second)
			foo(c)
		}()

		c.String(http.StatusOK, message)
	})

	r.GET("/panic", func(c *gin.Context) {
		f, err := os.Open("/Users/liangzuobin/Downloads/demo2.wmv")
		if err != nil {
			log.Fatal(err)
		}

		s, err := f.Stat()
		if err != nil {
			log.Fatal(err)
		}

		c.DataFromReader(http.StatusOK, s.Size(), "application/octet-stream", f, nil)
	})

	r.PUT("/upload", func(c *gin.Context) {
		body := c.Request.Body
		if body == nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		defer body.Close()

		handleReader(body)

		c.Status(http.StatusOK)
	})

	_ = r.Run(":8080")
}

func handleReader(r io.Reader) {
	obj := newreqobj(r)
	log.Printf("obj %#v", obj)
}

func newreqobj(r io.Reader) *reqobj {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Printf("new reqobj with err: %v", err)
		return nil
	}

	return &reqobj{
		ContentType:   "application/octet-stream",
		ContentLength: int64(len(b)),
		Body:          bytes.NewBuffer(b),
	}
}

type reqobj struct {
	ContentLength int64
	ContentType   string
	Body          io.Reader
}

const key = "key"

func foo(c *gin.Context) {
	val, _ := c.Get(key)
	if val == "" {
		panic("val is nil")
	}
}
