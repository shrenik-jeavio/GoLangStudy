// sum is a function that takes a slice of numbers and adds them together. what would its function signature look like in Go?

//func sum(xs []int) int
package main

import "fmt"


func main(){
	total := func(xs []int) int {
		t := 0
		for v := range xs {
			t += v
		}
		return t
	}
	x := []int{1,2,3,4,5,6,7,8,9,10,}
	fmt.Println(total(x)/len(x))

}
