package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	porta = flag.Int(
		"porta",
		8080,
		"Porta que o servidor HTTP irá escutar.")
)

func main() {
	flag.Parse()

	http.HandleFunc("/",
		func(w http.ResponseWriter,
			req *http.Request) {

			if req.Method == "GET" {
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte("[Echo] Só aceitamos requisições POST!"))
				return
			}

			// Ler o que o cliente enviou.
			msg, err := ioutil.ReadAll(req.Body)
			if err != nil {
				panic(err)
			}
			req.Body.Close()

			if len(msg) == 0 {
				w.Write([]byte("[Echo] Mensagem Vazia."))
				return
			}

			r := fmt.Sprintf("[Echo] %s", string(msg))
			w.Write([]byte(r))
		})

	// Escuta a passada como parâmetro esperando
	// por requições HTTP.
	http.ListenAndServe(
		fmt.Sprintf(":%d", *porta),
		nil)
}
