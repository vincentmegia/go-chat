package main

import (
	"log"
	"net"

	vmlog "vincentmegia.com/chat/pkg/log"
)

func main() {
	lw := vmlog.LogWriter{}
	lw.Init("DEBUG")

	log.Println("Starting server version:")
	listener, error := net.Listen("tcp", "localhost:4090")
	if error != nil {
		log.Println("Error listening for connections: ", error)
		return
	}
	defer listener.Close()

	log.Println("Server is listening to port 4090")
	for {
		connection, error := listener.Accept()
		if error != nil {
			log.Println("Error has occured accepting new connection: ", error)
		}
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	log.Println("Connection has been establised: ", connection.LocalAddr().String())
	defer connection.Close()

	// Read data sent by client
	buffer := make([]byte, 1024)

	for {
		n, error := connection.Read(buffer)
		if n == 0 {
			continue
		}
		if error != nil {
			log.Println("Byte receive error has occured: ", error)
			return
		}
		log.Printf("Received byte: %s", buffer[:n])
	}
}
