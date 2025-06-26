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

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("cotacao.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao abrir arquivo: %v\n", err)
		return
	}
	defer f.Close()

	// Verifica se o arquivo não está vazio
	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao verificar arquivo: %v\n", err)
		return
	}

	// Adiciona quebra de linha se arquivo não estiver vazio
	if fileInfo.Size() > 0 {
		if _, err := f.WriteString("\n"); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
			return
		}
	}

	// Adiciona timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	if _, err := f.WriteString(timestamp + " - "); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao escrever timestamp: %v\n", err)
		return
	}

	// Escreve os dados
	tamanho, err := f.Write(bodyBytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
		return
	}

	fmt.Printf("Dados adicionados com sucesso. Tamanho: %d bytes\n", tamanho)

}
