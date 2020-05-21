package main

type inter interface {
	Run(string) (bool, error)
}

type foo struct {
	Run func(string) (bool, error)
}

// var _ inter = &foo{} // compile error
