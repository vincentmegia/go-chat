package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("starting server version:")
	listener, error := net.Listen("tcp", "localhost:4090")
	if error != nil {
		fmt.Println("Error listening for connections: ", error)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening to port 4090")
	for {
		connection, error := listener.Accept()
		if error != nil {
			fmt.Println("Error has occured accepting new connection: ", error)
		}
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	fmt.Println("Connection has been establised: ", connection.LocalAddr().String())
	defer connection.Close()

	// Read data sent by client
	buffer := make([]byte, 1024)

	for {
		n, error := connection.Read(buffer)
		if n == 0 {
			continue
		}
		if error != nil {
			fmt.Println("Byte receive error has occured: ", error)
			return
		}
		fmt.Printf("Received byte: %s\n", buffer[:n])
	}
}
