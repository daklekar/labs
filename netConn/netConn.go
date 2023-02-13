package main

import(
	"fmt"
	"net"
)

func main(){
	for{
		listener, _ := net.Listen("tcp", ":8000")

		

		for{
			acc, _ := listener.Accept()

			x:=acc.RemoteAddr().String()
			a := acc.R
			fmt.Println(x,)

		}
	
	}
}