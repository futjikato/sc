package network

import (
	"net"
	"fmt"
	"github.com/futjikato/docker-sc/types"
	"encoding/json"
)

type Server struct {
	connection *net.UDPConn
	open bool
	MessageChannel chan types.StatSet
}

func (s *Server) Connect(port int) {
	conn, err := net.ListenUDP("udp4", getAddress(port))
	if err != nil {
		panic(err)
	}

	s.connection = conn

	fmt.Println(conn.LocalAddr().String())
}

func (s *Server) Listen() {
	s.open = true
	go func() {
		fmt.Println("Start listening")
		for s.open {
			payload := make([]byte, 1024)
			n, _, err := s.connection.ReadFromUDP(payload)

			if err != nil {
				fmt.Println(err)
				continue
			}

			if n <= 0 {
				continue
			}

			stats := types.StatSet{}
			err = json.Unmarshal(payload[:n], &stats)
			fmt.Println(string(payload[:n]))
			if err != nil {
				fmt.Println(err)
				continue
			}
			s.MessageChannel <- stats
		}
	}()
}

func (s *Server) Close() {
	s.connection.Close()
}

func getAddress(port int) (*net.UDPAddr) {
	addr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	return addr
}