package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"

	vmlog "vincentmegia.com/chat/pkg/log"
)

func main() {
	// set logging format
	lw := vmlog.LogWriter{}
	lw.Init("DEBUG")
	connection, error := net.Dial("tcp", "localhost:4090")
	if error != nil {
		log.Println("An error has occured establishing connection with server at localhost:4090")
		return
	}

	log.Println("Stupendousware chat has started and connected to server")
	log.Println("chat is ready")
	reader := bufio.NewReader(os.Stdin)
	for {
		line, error := reader.ReadBytes('\n')
		if error != nil {
			log.Println("an error has occured: ", error)
			continue
		}

		value := string(line[:])
		log.Println(value)
		if strings.Trim(value, "\r\n") == "exit" {
			log.Println("client is not exiting gracefully...")
			break
		}

		data := []byte(line)
		_, error = connection.Write(data)
		if error != nil {
			log.Println("Error has occured sending byte: ", error)
			continue
		}
	}
	defer connection.Close()
}
