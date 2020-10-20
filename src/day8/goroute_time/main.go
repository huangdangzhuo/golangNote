package main

import (
	"fmt"
	"runtime"
	"time"
)

func main()  {
	num:=runtime.NumCPU()
	runtime.GOMAXPROCS(num-1)
	for i:=0;i<1000 ;i++ {
		go func() {
			for {

				select {
				case <-time.After(time.Microsecond):
					fmt.Println("get data timeout")
				}
			}
		}()
	}

	time.Sleep(time.Second*100000)

}



