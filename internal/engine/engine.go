package engine

import (
	"fmt"
	"strconv"

	"github.com/Saad7890-web/neurocache/internal/cluster"
	"github.com/Saad7890-web/neurocache/internal/persistence"
	"github.com/Saad7890-web/neurocache/internal/protocol"
)


type Engine struct {
	shards *cluster.ShardManager
	aof *persistence.AOF
}

func NewEngine() *Engine {
	aof, _ := persistence.NewAOF("data.aof")
	return &Engine{
		shards: cluster.NewShardManager(8),
		aof: aof,
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
	e.aof.Write(fmt.Sprintf("SET %s %s", key, value))
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
		e.aof.Write(fmt.Sprintf("DEL %s", cmd.Args[0]))
		ok := e.shards.Del(cmd.Args[0])

		if ok {
			return ":1\r\n"
		}

		return ":0\r\n"

	default:
		return "-ERR unknown command\r\n"
	}
}