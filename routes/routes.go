package routes

import (
	"log"
	"net/http"

	"github.com/cyph3rk/cotacao_dolar/config"
	"github.com/cyph3rk/cotacao_dolar/controller"
)

func HandleResquest() {

	cfg := config.Get()

	log.Printf("Starting server on %s", cfg.Server.Port)

	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/cotacao", controller.PegaCotacao)
	http.ListenAndServe(cfg.Server.Port, nil)
}
