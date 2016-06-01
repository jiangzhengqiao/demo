package main

import (
	"errors"
	"github.com/hprose/hprose-go"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

type myService struct{}

func (myService) Swap(a int, b int) (int, int) {
	return b, a
}

func (myService) Sum(args ...int) (int, error) {
	if len(args) < 2 {
		return 0, errors.New("Requires at least two parameters")
	}
	a := args[0]
	for i := 1; i < len(args); i++ {
		a += args[i]
	}
	return a, nil
}

func main() {
	server := hprose.NewTcpServer("tcp://192.168.21.3:1234/")
	server.AddFunction("hello", hello)
	server.AddMethods(myService{})
	server.Start()
}
