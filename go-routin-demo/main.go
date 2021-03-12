package main

import (
	"fmt"
	"time"
)

func printString(s string) {

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, ":", i)
	}

}

func main() {

	go printString("Call using go")

	printString("Calling Direct (without go)")

	fmt.Println("Done.")

}
