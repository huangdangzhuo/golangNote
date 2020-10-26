package main

import (
	"fmt"
)
// Strategy 策略类接口
type Strategy interface {
	DoDperation(num1,num2 int) int
}

// OperationAdd 加法策略类
type OperationAdd struct {
}

// NewOperationAdd 实例化加法策略类
func NewOperationAdd() *OperationAdd {
	return &OperationAdd{}
}

// DoOperationAdd
func (c *Circle) Draw() {
	fmt.Println("Circle Draw method.")
}

// RedShapeDecorator 红色装饰器
type RedShapeDecorator struct {
	DecoratedShape Shape
}

// NewRedShapeDecorator 实例化红色装饰器
func NewRedShapeDecorator(decShape Shape) *RedShapeDecorator {
	return &RedShapeDecorator{DecoratedShape: decShape,
	}
}

// SetRedBorder 装饰器方法
func (rs *RedShapeDecorator) SetRedBorder()  {
	fmt.Println("Border Color:red")
}

// Draw 实现Shape接口的方法
func (rs *RedShapeDecorator) Draw()  {
	rs.DecoratedShape.Draw()
	rs.SetRedBorder()
}