package persistence

import (
	"os"
	"sync"
)


type AOF struct {
	file *os.File
	mu sync.Mutex
}

func NewAOF(path string) (*AOF, error) {
	file,err := os.OpenFile(
		path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return nil, err
	}
	return &AOF{file: file}, nil
}

func (a *AOF) Write(command string) {

	a.mu.Lock()
	defer a.mu.Unlock()

	a.file.WriteString(command + "\n")
}