
package main


import (
	"fmt"
	"math"
	"image/color"	
)

type Point struct{X, Y float64}



type ColoredPoint struct {
	*Point
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

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	var p = ColoredPoint{&Point{1, 1}, red}
	var q = ColoredPoint{&Point{5, 4}, blue}

	fmt.Println(p.Distance(*q.Point))	
}
