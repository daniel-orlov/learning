package main

import (
	"bufio"
	"fmt"
	es "github.com/pkg/errors"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		err = es.Wrap(err, "Listening failed: tcp, :8080")
		fmt.Println(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			err = es.Wrap(err, "Listener Accept() failed")
			fmt.Println(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	name := conn.RemoteAddr().String()

	fmt.Printf("%+v connected\n", name)
	conn.Write([]byte("Hello, " + name + "\n\r"))

	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "Exit" || text == "exit" {
			conn.Write([]byte("Bye\n\r"))
			fmt.Println(name, "disconnected")
			break
		} else if text != "" {
			fmt.Println(name, "says", text)
			conn.Write([]byte("You entered " + text + "\n\r"))
		}
	}
}
