package main

import (
	"testing"
	"time"
)

func TestStudent_Save(t *testing.T) {
	stu :=student{
		Name: "stu10",
		Sex:  "man",
		Age:  "18",
	}
	err := stu.Save()
	if err!=nil {
		t.Fatalf("save student error:%v",err)
	}
}

func TestStudent_Load(t *testing.T) {
	stu :=student{
		Name: "stu10",
		Sex:  "man",
		Age:  "18",
	}
	err := stu.Save()
	if err!=nil {
		t.Fatalf("save student error:%v",err)
	}
	time.Sleep(time.Second*10)
	stu1 :=student{}
	err1 := stu1.Load()
	if err1!=nil {
		t.Fatalf("load student error:%v",err1)
	}
	if stu.Name!=stu1.Name {
		t.Fatalf("load student Name error")
	}
	if stu.Sex!=stu1.Sex {
		t.Fatalf("load student Sex error")
	}
	if stu.Age!=stu1.Age {
		t.Fatalf("load student Age error")
	}
}