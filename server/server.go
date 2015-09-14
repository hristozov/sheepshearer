package server

import (
	"fmt"
	"net"
	"os"
)

// Start just starts the listener
func Start() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s\n", err.Error())
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error %s\n", err.Error())
			return
		}
		go handleConnection(conn)
	}
}
