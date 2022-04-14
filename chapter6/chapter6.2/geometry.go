package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// Distance 函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) ChangeTest() {
	p.X = 0
	p.Y = 0
}

func (p *Point) Change() {
	p.X = 0
	p.Y = 0
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p := Point{1, 1}
	p.Change()
	fmt.Println(p)

	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	p = Point{1, 2}
	p.ScaleBy(3)
	fmt.Println(p)
}
