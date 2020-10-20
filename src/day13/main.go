package main

import (
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"
)

func test() {
	var mut sync.Mutex
	maxSize := 10
	counter := 0

	// 排水口
	go func() {
		for {
			mut.Lock()
			if counter == maxSize {
				for i := 0; i < maxSize; i++ {
					counter--
					log.Printf(" OUTPUT counter = %d", counter)
				}
			}
			mut.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()

	// 注水口
	for {
		mut.Lock()
		if counter == 0 {
			for i := 0; i < maxSize; i++ {
				counter++
				log.Printf(" INPUT counter = %d", counter)
			}
		}
		mut.Unlock()
		time.Sleep(1 * time.Second)
	}
}

//func main()  {
//	s:= Student{
//		Name:    "dskk",
//		Age:     20,
//		Address: "dskfaskjn",
//	}
//	Test1(s)
//}

///////////////////////// 反射 /////////////////////////
// Student 学生结构体
type Student struct {
	Name    string
	Age     int
	Address string
}

// Say Student结构体的方法
func (s Student) Say(msg string) {
	fmt.Println(msg)
}

// PrintInfo Student结构体的方法
func (s Student) PrintInfo() {
	fmt.Printf("姓名:%s\t年龄:%d\t地址:%s\n", s.Name, s.Age, s.Address)
}

func Test1(i interface{}) {
	// 获取i的类型
	rType := reflect.TypeOf(i)
	fmt.Println("i的类型是:", rType.Name()) // Student
	fmt.Println("i的种类是:", rType.Kind()) // struct

	// 获取i的字段信息
	rValue := reflect.ValueOf(i)
	fmt.Println("i的值是:", rValue)

	// 获取i的字段信息
	for i := 0; i < rValue.NumField(); i++ {
		filed := rType.Field(i)
		value := rValue.Field(i).Interface()
		fmt.Printf("字段名称:%s,字段类型:%s,字段值:%v\n", filed.Name, filed.Type, value)

	}
	// 获取i的方法信息
	for i := 0; i < rValue.NumMethod(); i++ {
		method := rType.Method(i)
		fmt.Printf("方法的名称:%s,方法的类型:%s\n", method.Name, method.Type)
	}

}

// Shape 模型接口
type Shape interface {
	Draw()
}

// Circle 圆类型
type Circle struct {
}

// NewCircle 实例化圆形
func NewCircle() *Circle {
	return &Circle{}
}

// Draw 输出方法,实现Shape接口
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