package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	buf := make([]byte, 1024)
	conn, err := net.Dial("udp", "255.255.255.255:3000")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")
	_, err = bufio.NewReader(conn).Read(buf)
	if err == nil {
		fmt.Printf("%s\n", string(buf))
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}
