package main

import (
	"fmt"
	"math"
)

// 定义一个基本的几何形状的接口
type geometry interface {
	area() float64 // 面积

	perim() float64 // 周长
}

// 长方体与圆形实体
type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

// 长方体实现接口
func (s square) area() float64 {
	return s.height * s.width
}
func (s square) perim() float64 {
	return (s.height + s.width) * 2
}

// 圆形实现接口
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("面积：", g.area())
	fmt.Println("周长：", g.perim())
}

func main() {
	s := square{width: 3, height: 4}
	c := circle{radius: 8}
	measure(c)
	measure(s)
}
