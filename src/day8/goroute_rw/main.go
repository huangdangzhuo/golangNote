package main

import (
	"fmt"
	"time"
)

func main()  {
	intChan :=make(chan int,10)
	intChan1 :=make(chan int,10)

	go func() {
		for i:=0;i<10 ;i++  {
			intChan<-i
			intChan1<-i*i
		}
	}()
	for  {
		select {
		case v:= <-intChan:
			fmt.Println(v)
		case v:= <-intChan1:
			fmt.Println(v)
		default:
			fmt.Println("get data timeout")
			time.Sleep(time.Second)

		}
	}


}



