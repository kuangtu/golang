package main

import "fmt"

type Polygon struct {
	Sides int
}

func (p *Polygon) NSides() int {
	return p.Sides
}

//继承了匿名结构体的方法Nsides()
type Triangle struct {
	Polygon
}

func main() {
	t := Triangle{Polygon{Sides: 3}}

	fmt.Println(t.NSides())
}
