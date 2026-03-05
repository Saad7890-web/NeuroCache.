package protocol

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

func ParseCommand(reader *bufio.Reader) (*Command, error){
	line, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	line = strings.TrimSpace(line)
	if len(line) == 0 || line[0] != '*' {
		return nil, errors.New("Invalid Resp array")
	}

	argCount, err := strconv.Atoi(line[1:])

	if err != nil {
		return nil, err
	}

	args := make([]string, 0, argCount)

	for i := 0; i < argCount; i++ {

		lenLine, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		lenLine = strings.TrimSpace(lenLine)

		if lenLine[0] != '$' {
			return nil, errors.New("invalid bulk string")
		}

		strLen, err := strconv.Atoi(lenLine[1:])
		if err != nil {
			return nil, err
		}

		data := make([]byte, strLen+2)
		_, err = reader.Read(data)
		if err != nil {
			return nil, err
		}

		arg := strings.TrimSpace(string(data))
		args = append(args, arg)
	}

	cmd := &Command{
		Name: strings.ToUpper(args[0]),
		Args: args[1:],
	}

	return cmd, nil
}