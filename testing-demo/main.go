package main
import "fmt"
func Calculate(x int) (result int) {
	result = x + 2
	return
}
func main() {
	fmt.Println("Welcome to go testing")
	result := Calculate(2)
	fmt.Println(result)
}