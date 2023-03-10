package main

import (
	"fmt"
	"io"
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

	// c := make(chan string)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		defer conn.Close()
		go handleConnection(conn)
		// fmt.Println(<- c)
	}
}

func handleConnection(conn net.Conn) {
	// cTraffic := make(chan string)
	
	for {
		buf := make([]byte, 4096)
		_, err := conn.Read(buf)

		if err != nil {
			if err == io.EOF {
				fmt.Println("has reached the end of file.")
				break
			} else {
				fmt.Println("The read error is: ", err.Error())
				os.Exit(1)
			}
		}

		fmt.Println("The content is: ", string(buf))
		_, myError := conn.Write([]byte("+PONG\r\n"))
		if myError != nil {
			fmt.Println("The write error is: ", myError.Error())
			os.Exit(1)
		}
		
		// c <- "one connection handled"
	}


	// go handleTraffic(content, conn, cTraffic)

	// <- cTraffic
	// fmt.Println(<- cTraffic)
	// c <- "one connection handled"

}

// func handleTraffic(content []byte, conn net.Conn, cTraffic chan string) {
// 	// fmt.Println("The content is: ", string(content))

// 	_, myError := conn.Write([]byte("+PONG\r\n"))
// 	if myError != nil {
// 		fmt.Println("The write error is: ", myError.Error())
// 		os.Exit(1)
// 	}

// 	cTraffic <- "traffic handled"
// }
