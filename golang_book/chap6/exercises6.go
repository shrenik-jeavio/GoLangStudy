
/* write function to swap two integers */
package main
import "fmt"

func swap(x, y *int){
	*x, *y = *y, *x
}
func main(){
	x := 10
	y := 15
	fmt.Println("x:", x, "| y:", y)
	swap(&x, &y)
	fmt.Println("after swap")
	fmt.Println("x:", x, "| y:", y)
}