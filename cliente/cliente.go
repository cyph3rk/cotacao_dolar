package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Cliente...")

	ConsultaCotacaoDolarNoServer()

}

func ConsultaCotacaoDolarNoServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	//io.Copy(os.Stdout, res.Body)

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	tamanho, err := f.Write(bodyBytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso. Tamanho: %d bytes\n", tamanho)
	f.Close()

}
