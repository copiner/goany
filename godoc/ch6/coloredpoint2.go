
package main


import (
	"fmt"
	"math"
	"image/color"	
)

type Point struct{X, Y float64}



type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 { //p -- receiver
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64){
	p.X *= factor
	p.Y *= factor	
}


func main(){

	p := Point{1, 2}
	q := Point{4, 6}	

	distanceFromP := p.Distance //method value
	fmt.Println(distanceFromP(q))

	var origin Point
	fmt.Println(distanceFromP(origin))

	scaleP := p.ScaleBy
	scaleP(2)
	fmt.Println(p)

	m := Point{1, 2}
	n := Point{4, 6}	
	
	distance := Point.Distance //method expression
	fmt.Println(distance(m, n))

	scale := (*Point).ScaleBy
	scale(&m, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)	
}
