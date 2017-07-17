package main

import ("fmt"
	"math"
)
//Circle
type Circle struct{
	x, y, r float64
}

func circleArea(c *Circle) float64{
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64{
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64{
	return 2 * math.Pi * c.r
}

//Rectangle
type Rectangle struct{
	x1, y1, x2, y2 float64
}

func distance(x1, x2, y1, y2 float64) float64{
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64{
	l := distance(r.x1, r.x2, r.y1, r.y2)
	w := distance(r.x1, r.x2, r.y1, r.y2)
	return  l * w
}

func (r *Rectangle) perimeter() float64{
	l := distance(r.x1, r.x2, r.y1, r.y2)
	w := distance(r.x1, r.x2, r.y1, r.y2)
	return  2 * (l + w)
}

//interface
type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64{
	var area float64
	for _, s := range shapes{
		area += s.area()
	}
	return area
}

func main(){
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}
	fmt.Println("Circle Area")
	//fmt.Println(circleArea(&c))
	fmt.Println(c.area())

	fmt.Println()

	fmt.Println("Rectangle Area")
	fmt.Println(r.area())

	fmt.Println()

	fmt.Println("Total Area")
	fmt.Println(totalArea(&c, &r))

	fmt.Println()
	fmt.Println("Circle Perimeter Area")
	fmt.Println(c.perimeter())

	fmt.Println()
	fmt.Println("Rectangle Perimeter Area")
	fmt.Println(r.perimeter())

}
