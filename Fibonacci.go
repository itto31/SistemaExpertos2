package main

import "fmt"

func fibonaccia() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonaccia()
	for i := 0; i <= 20; i++ {
		fmt.Printf("fibonacci(%d): %d\n", i, f())
	}
}
