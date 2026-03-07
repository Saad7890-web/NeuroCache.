package engine

import (
	"fmt"
	"strconv"

	"github.com/Saad7890-web/neurocache/internal/cluster"
	"github.com/Saad7890-web/neurocache/internal/protocol"
)


type Engine struct {
	shards *cluster.ShardManager
}

func NewEngine() *Engine {

	return &Engine{
		shards: cluster.NewShardManager(8),
	}
}


func (e *Engine) Execute(cmd *protocol.Command) string {

	switch cmd.Name {

	case "SET":
		if len(cmd.Args) < 2 {
		return "-ERR wrong number of arguments\r\n"
	}

	key := cmd.Args[0]
	value := cmd.Args[1]

	ttl := 0

	if len(cmd.Args) == 4 && cmd.Args[2] == "EX" {
		t, err := strconv.Atoi(cmd.Args[3])
		if err == nil {
			ttl = t
		}
	}

	e.shards.Set(key, value, ttl)

	return "+OK\r\n"
	case "GET":
		if len(cmd.Args) != 1 {
			return "-ERR wrong number of arguments\r\n"
		}

		val, ok := e.shards.Get(cmd.Args[0])
		if !ok {
			return "$-1\r\n"
		}

		return fmt.Sprintf("$%d\r\n%s\r\n", len(val), val)

	case "DEL":
		if len(cmd.Args) != 1 {
			return "-ERR wrong number of arguments\r\n"
		}

		ok := e.shards.Del(cmd.Args[0])

		if ok {
			return ":1\r\n"
		}

		return ":0\r\n"

	default:
		return "-ERR unknown command\r\n"
	}
}