package facade

import (
	"fmt"
)

// Shape 模型接口
type Shape interface {
	Draw()
}

// Circle 圆类型
type Circle struct {
}

// Rectangle 矩形类
type Rectangle struct {
}

// Square 矩形类
type Square struct {
}

// NewCircle 实例化圆形类
func NewCircle() *Circle {
	return &Circle{}
}

// Draw 圆形类实现Shape接口
func (c *Circle) Draw() {
	fmt.Println("Circle Draw method.")
}

// NewRectangle 实例化矩形类
func NewRectangle() *Rectangle {
	return &Rectangle{}
}

// Draw 矩形类实现Shape接口
func (r *Rectangle) Draw() {
	fmt.Println("Rectangle Draw method.")
}

// NewSquare 实例化正方形类
func NewSquare() *Square {
	return &Square{}
}

// Draw 正方形类实现Shape接口
func (s *Square) Draw() {
	fmt.Println("Square Draw method.")
}

// ShapeMaker 外观类
type ShapeMaker struct {
	circle Circle
	square Square
	rectangle Rectangle
}

// NewShapeMaker 实例化外观类
func NewShapeMaker() *ShapeMaker {
	return &ShapeMaker{
		circle: Circle{},
		square: Square{},
		rectangle: Rectangle{},
	}
}

// DrawCircle 实现Circle的Draw方法
func (sm *ShapeMaker) DrawCircle() {
	sm.circle.Draw()
}

// DrawRectangle 实现Rectangle的Draw方法
func (sm *ShapeMaker) DrawRectangle() {
	sm.rectangle.Draw()
}

// DrawSquare 实现Square的Draw方法
func (sm *ShapeMaker) DrawSquare() {
	sm.square.Draw()
}
