// Fibonacci Sequence
package main

import "fmt"

func Fibonacci(n int) {
	a := 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, "")
	}
	fmt.Println()
}
func main() {
	var n int
	fmt.Println("value of n: ")
	fmt.Sacn(&n)
	fmt.Printf("Fibonacci Sequence up to %d is: ", n)
	Fibonacci(n)
}
