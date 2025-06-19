package routes

import (
	"net/http"

	"github.com/cyph3rk/cotacao_dolar/controller"
)

func HandleResquest() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/cotacao", controller.PegaCotacao)
	http.ListenAndServe(":8080", nil)
}
