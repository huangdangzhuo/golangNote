package main

import (
	"fmt"
	"net/http"
	"time"
)


var url =[]string{
	"https://www.baidu.com",
	"https://www.google.com",
	"https://www.taobao.com",
}


func main()  {
	for _,v:=range url {
		c:=http.Client{
			Timeout:       time.Second*2,
		}

		res, err := c.Head(v)
		if err != nil {
			fmt.Printf("head %s failed, err:%v\n",v, err)
			return
		}
		fmt.Println("head succ,status:%v\n",res.Status)
	}
}