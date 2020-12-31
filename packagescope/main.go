package main

import "fmt"

var declareTwice = 10

func nested() {
	var declareTwice = 5
	fmt.Println("Inside the nested function: ", declareTwice)
}

func main() {
	fmt.Println("Inside the main function: ", declareTwice)
	nested()
}
