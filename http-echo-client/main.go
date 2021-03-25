package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	servidor = flag.String(
		"servidor",
		"localhost:8080",
		"Host:Porta do servidor HTTP")

	mensagem = flag.String(
		"mensagem",
		"Olá",
		"Mensagem a ser enviada para o servidor")
)

func main() {
	flag.Parse()

	// Executa requisição HTTP POST.
	resp, err := http.Post(*servidor,
		"text/plain",
		bytes.NewBufferString(*mensagem))
	if err != nil {
		panic(err)
	}

	// Ler o que o servidor enviou.
	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()

	fmt.Printf(
		"Mensagem do servidor %s: %s\n",
		*servidor,
		string(msg))
}
