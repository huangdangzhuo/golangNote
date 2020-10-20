package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main()  {
	conn,err := net.Dial("tcp","localhost:50000")
	if err != nil {
		fmt.Println("Dial failed err:",err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input,_ := inputReader.ReadString('\n')
		trimedInput :=strings.Trim(input,"\r\n")
		if trimedInput == "Q" {
			return
		}
		_,err := conn.Write([]byte(trimedInput))
		if err!=nil {
			return
		}
	}
}
