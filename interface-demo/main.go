package main

import (
	"fmt"
)

type i interface {
	Get() int
	Put(v int)
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

type S struct {
	i int
}

func (s *S) Get() int {
	return s.i
}

func (s *S) Put(v int) {
	s.i = v
}

func main() {
	fmt.Println("Welcome to inteface demo")
	var strcutVar S
	strcutVar.Put(10)
	fmt.Println(strcutVar.Get())

	var r rect = rect{3, 4}
	fmt.Println(r.area())
}
