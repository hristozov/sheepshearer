package server

import (
	"fmt"
	"net"
	"os"
	"sheepshearer/parser"
)

const bufferSize = 1024

func handleConnection(conn net.Conn) {
	//conn.SetDeadline(time.Now().Add(5000))
	fmt.Println("Opened connection from ", conn.RemoteAddr())

	for {
		var requestBuffer = make([]byte, bufferSize)
		readlen, err := conn.Read(requestBuffer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error %s\n", err.Error())
			conn.Close()
			return
		}

		if readlen == 0 {
			fmt.Printf("Connection closed by remote host\n")
			conn.Close()
			return
		}

		requestRaw := string(requestBuffer[:readlen])
		fmt.Println(requestRaw)

		var response string

		req, error := parser.Parse(requestRaw)
		if error != nil {
			fmt.Fprintf(os.Stderr, "Error %s\n", error.Error())
			response = BuildErrorResponse(400)
		} else {
			response = BuildOkResponse([]byte("Hello "+req.Path), "text/html")
		}

		conn.Write([]byte(response))
		conn.Close()
		break
	}
}
