package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cyph3rk/cotacao_dolar/service"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func PegaCotacao(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
		http.Error(w, "Request cancelada pelo cliente", http.StatusRequestTimeout)
		return
	case <-time.After(5 * time.Second):
		cotacao, err := service.PegaCotacao(ctx)
		if err != nil {
			log.Println("Erro ao buscar cotação:", err)
			http.Error(w, "Erro ao buscar cotação", http.StatusInternalServerError)
			return
		}

		resposta := map[string]string{
			"cotacao": cotacao.Bid,
		}

		jsonResposta, err := json.Marshal(resposta)
		if err != nil {
			log.Println("Erro ao converter resposta para JSON:", err)
			http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResposta)
		log.Println("Cotação enviada com sucesso")
	}
}
