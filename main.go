package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/cyph3rk/cotacao_dolar/database"
	"github.com/cyph3rk/cotacao_dolar/service"
)

func main() {
	fmt.Println("Olá mundo")

	database.ConectaComBancoDeDados()

	ctx := context.Background() // Contexto raiz
	cotacao, err := service.PegaCotacao(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Println("Timeout ao buscar cotação")
		} else {
			log.Printf("Erro: %v", err)
		}
		return
	}
	fmt.Println(cotacao)

}
