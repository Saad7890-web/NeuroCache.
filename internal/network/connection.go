package network

import (
	"bufio"
	"net"
)

func handleConnection(conn net.Conn){
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		data, err := reader.ReadBytes('\n')
		if err != nil {
			return
		}

		conn.Write([]byte("OK\n"))
		_ = data
	}
}