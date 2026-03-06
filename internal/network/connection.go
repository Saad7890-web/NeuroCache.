package network

import (
	"bufio"
	"net"

	"github.com/Saad7890-web/neurocache/internal/engine"
	"github.com/Saad7890-web/neurocache/internal/protocol"
)

func handleConnection(conn net.Conn, eng *engine.Engine) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {

		cmd, err := protocol.ParseCommand(reader)
		if err != nil {
			return
		}

		response := eng.Execute(cmd)

		conn.Write([]byte(response))
	}
}