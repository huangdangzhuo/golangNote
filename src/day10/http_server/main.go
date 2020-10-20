package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter,r *http.Request){
	fmt.Println("handle hello")
	fmt.Fprintf(w,"hello")
}
func user(w http.ResponseWriter,r *http.Request){
	fmt.Println("handle user")
	fmt.Fprintf(w,"user")
}

func main()  {
	http.HandleFunc("/",Hello)
	http.HandleFunc("/user",user)
	err:=http.ListenAndServe("0.0.0.0:8880",nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}