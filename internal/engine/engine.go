package engine

import (
	"fmt"

	"github.com/Saad7890-web/neurocache/internal/kv"
	"github.com/Saad7890-web/neurocache/internal/protocol"
)


type Engine struct {
	store *kv.Store
}

func NewEngine() *Engine {
	return &Engine{
		store: kv.NewStore(),
	}
}


func (e *Engine) Execute(cmd *protocol.Command) string {

	switch cmd.Name {

	case "SET":
		if len(cmd.Args) != 2 {
			return "-ERR wrong number of arguments\r\n"
		}

		e.store.Set(cmd.Args[0], cmd.Args[1])
		return "+OK\r\n"

	case "GET":
		if len(cmd.Args) != 1 {
			return "-ERR wrong number of arguments\r\n"
		}

		val, ok := e.store.Get(cmd.Args[0])
		if !ok {
			return "$-1\r\n"
		}

		return fmt.Sprintf("$%d\r\n%s\r\n", len(val), val)

	case "DEL":
		if len(cmd.Args) != 1 {
			return "-ERR wrong number of arguments\r\n"
		}

		ok := e.store.Del(cmd.Args[0])

		if ok {
			return ":1\r\n"
		}

		return ":0\r\n"

	default:
		return "-ERR unknown command\r\n"
	}
}