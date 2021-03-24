package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

var port = flag.String("server", "127.0.0.1:3000", "Host:Porta do servidor UDP")
var msg = flag.String("msg", "Olá", "Mensagem a ser enviada para o servidor")

func main() {
	flag.Parse()
	

	// Abrindo conexão
	conn, err := net.Dial("udp", *port)
	if err != nil {
		fmt.Printf("Erro: %v", err)
		return
	}
	defer conn.Close()

	// Enviando mensagem.
	fmt.Fprintf(conn, *msg)

	// Resposta da conexão.
	buf := make([]byte, 1024)
	_, err = bufio.NewReader(conn).Read(buf)
	if err == nil {
		fmt.Printf("%s\n", string(buf))
	} else {
		fmt.Printf("Erro: %v\n", err)
	}
	conn.Close()
}
