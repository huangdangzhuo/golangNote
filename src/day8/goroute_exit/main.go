package main

import (
	"fmt"
)



func write(ch chan int,exitChan chan struct{}){
	for i:=0;i<10 ;i++  {
		ch<-i
		fmt.Println("put data:",i)
	}
	exitChan <-struct{}{}
}
func read(ch chan int,exitChan chan struct{}){
	for{
		v,ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	exitChan <-struct{}{}
}

func main()  {
	intChan :=make(chan int,10)
	exitChan := make(chan struct{},2)
	go write(intChan,exitChan)
	go read(intChan,exitChan)

	<-exitChan

}



