package main

import (
	"fmt"
	"log"
	"net"
)

const bufSize = 1024

func main() {
	s, err := net.ListenPacket("udp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	for {
		buf := make([]byte, bufSize)
		n, addr, err := s.ReadFrom(buf)
		if err != nil {
			continue
		}
		fmt.Printf("Chegou mensagen de %s: %s\n", addr.String(), buf[:n])

		reply := fmt.Sprintf("[Server Echo] %s", string(buf[:n]))
		s.WriteTo([]byte(reply), addr)
	}
}
