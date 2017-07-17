// write fib seq is defined as fib(0) = 0, fib(1) = 1
package main

import "fmt"

func fib(x uint) uint{
	/*
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return fib(x-1) + factorial(x-2)
	*/
	switch x {
		case 0:
			return 0
		case 1:
			return 1
		default:
			return fib(x-1) + fib(x-2)

	}
}
func main(){
	fmt.Println(fib(16))
}