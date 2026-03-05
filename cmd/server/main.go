package main

import (
	"log"

	"github.com/Saad7890-web/neurocache/internal/network"
)

func main() {
	server := network.NewServer(":6381")

	log.Println("NeuroCache server starting on port 6381")

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}