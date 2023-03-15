package main

import (
	"net"
	"fmt"
	"bufio"
)

func main(){
	data_stream, e := net.Listen("tcp", ":8080")
	if e != nil{
		fmt.Println(e)
		return
	}
	defer data_stream.Close()

	for{
		connection, e := data_stream.Accept()
		if e != nil{
			fmt.Println(e)
			return
		}
		
		go handle_con(connection)
	}
}

func handle_con(connection net.Conn){
	for{
		data, e := bufio.NewReader(connection).ReadString('\n')
		if e != nil{
			fmt.Println(e)
			return
		}
		fmt.Println(data)
		
	}
	connection.Close()
}