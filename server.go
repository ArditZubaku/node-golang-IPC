package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const socketPath = "/tmp/ipc.sock"

func main() {
	// Clean old socket
	_ = os.Remove(socketPath)

	ln, err := net.Listen("unix", socketPath)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Println("Go server listening on", socketPath)

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewScanner(conn)

	for reader.Scan() {
		msg := reader.Text()
		fmt.Println("Received from Node:", msg)

		resp := "Go received: " + msg + "\n"
		_, _ = conn.Write([]byte(resp))
	}
}
