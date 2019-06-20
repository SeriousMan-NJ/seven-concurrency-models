package main

import (
	"bufio"
	"fmt"
	"net"
)

func ConnectionHandler(socket net.Conn) {
	in := bufio.NewReader(socket)
	out := bufio.NewWriter(socket)

	data := make([]byte, 1024)

	defer socket.Close()

	for {
		n, err := in.Read(data)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print(string(data[:n]))

		_, err = out.Write(data[:n])
		out.Flush()

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	server, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer server.Close()

	for {
		socket, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer socket.Close()

		go ConnectionHandler(socket)
	}
}
