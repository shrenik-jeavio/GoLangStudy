
// func takes an int & halves ur & return true if it was even or false if it was odd
package main

import "fmt"


func half(x int) (int, bool) {
	return x/2, x%2==0
}

func main(){

	y := 21
	fmt.Println(half(y))
}
