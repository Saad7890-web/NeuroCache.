package network

import (
	"bufio"
	"net"

	"github.com/Saad7890-web/neurocache/internal/protocol"
)

func handleConnection(conn net.Conn){
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {

		
		cmd, err := protocol.ParseCommand(reader)
		if err != nil {
			return
		}

		conn.Write([]byte("OK\n"))

		_ = cmd
	}
}