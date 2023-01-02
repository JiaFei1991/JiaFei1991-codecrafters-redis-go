package main

import (
	"fmt"
	"net"
	"os"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage hell yeah
	
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	// _, err = l.Accept()
	// if err != nil {
	// 	fmt.Println("Error accepting connection: ", err.Error())
	// 	os.Exit(1)
	// }

	c := make(chan string)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn, c)
	}

	fmt.Println(<- c)
}

func handleConnection(conn net.Conn, c chan string) {
	content := []byte{}

	length, err := conn.Read(content)

	if err != nil {
		fmt.Println("The read error is: ", err.Error())
	}

	fmt.Println("Read content length is: ", length)
	fmt.Println("The content is: ", string(content))

	c <- "traffic handled"
}
