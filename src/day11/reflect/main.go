package main

import (
	"fmt"
	"reflect"
)

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