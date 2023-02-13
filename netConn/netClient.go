
package main

import(
	//"fmt"
	"net"
)

func main(){
	//conn := net.Conn()
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
	conn.Write([]byte{8})
	conn.Close()
}