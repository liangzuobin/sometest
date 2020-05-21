package main

import "log"

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
