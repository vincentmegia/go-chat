package main

import (
	"fmt"
	"net"
)

func main() {
	connection, error := net.Dial("tcp", "localhost:4090")
	if error != nil {
		fmt.Println("An error has occured establishing connection with server at localhost:4090")
		return
	}

	data := []byte("Hello server!")
	_, error = connection.Write(data)
	if error != nil {
		fmt.Println("Error has occured sending byte: ", error)
		return
	}
	defer connection.Close()
}
