package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Connecting to server...")

	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("Connection succeeded")
	// send some data over our connection
	// fmt.Fprintf(conn, "Hi there!\n")
	// status, err := bufio.NewReader(conn).ReadString('\n')
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(status)

	go handleReturnMessages(conn)

	go userInputHandle(conn)

	// defer conn.Close()

}

// Setting up this function to handle input from the user in terminal
func userInputHandle(c net.Conn) {
	defer c.Close()

	// Setup reader to take user input from STDIN
	reader := bufio.NewReader(os.Stdin)
	// get input from user
	for {
		fmt.Print("me: ")
		text, _ := reader.ReadString('\n')

		// send text to server
		if _, err := io.WriteString(c, text); err != nil {
			fmt.Printf("Error: %v", err)
		}
	}

}

func handleReturnMessages(c net.Conn) {
	defer c.Close()

	for {
		// setup reader and return bytes
		serverResponse, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		// setup a writer to stdout
		if _, err := bufio.NewWriter(os.Stdout).WriteString(serverResponse); err != nil {
			log.Fatalf("Error: %v", err)
		}
		// fmt.Println(serverResponse)

	}

}
