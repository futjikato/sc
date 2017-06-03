package main

import (
	"github.com/futjikato/docker-sc/network"
	"fmt"
	"flag"
	"github.com/futjikato/docker-sc/types"
)

func main() {
	port := flag.Int("port", 41825, "Port to listen for UDP data")
	flag.Parse()

	mc := make(chan types.StatSet)

	s := network.Server{MessageChannel: mc}
	s.Connect(*port)
	s.Listen()
	defer s.Close()

	for {
		message := <-mc
		fmt.Println(message.Ts)
	}
}