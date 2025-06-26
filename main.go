package main

import (
	"fmt"
	"log"

	"github.com/cyph3rk/cotacao_dolar/config"
	"github.com/cyph3rk/cotacao_dolar/database"
	"github.com/cyph3rk/cotacao_dolar/routes"
)

func main() {
	fmt.Println("Servidor...")

	if err := config.LoadConfig("config/config.json"); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	database.ConectaComBancoDeDados()

	routes.HandleResquest()

	// ctx := context.Background()
	// cotacao, err := service.PegaCotacao(ctx)
	// if err != nil {
	// 	if errors.Is(err, context.DeadlineExceeded) {
	// 		log.Println("Timeout ao buscar cotação")
	// 	} else {
	// 		log.Printf("Erro: %v", err)
	// 	}
	// 	return
	// }
	// fmt.Println(cotacao)

}
