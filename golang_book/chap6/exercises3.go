// write funcation that take one variadic oara that finds the freatest number in a list of numbers

package main

import "fmt"

func max(xs ...int) int{
	var max int
	for i, x := range xs{
		if i == 0 || x > max{
			max = x
		}
	}
	return max
}

func main(){
	xs := []int{1,2,3}
	fmt.Println(max(xs...))
}