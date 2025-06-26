package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/cyph3rk/cotacao_dolar/config"
	"github.com/cyph3rk/cotacao_dolar/database"
	"github.com/cyph3rk/cotacao_dolar/models"
)

func PegaCotacao(ctx context.Context) (models.Cotacao, error) {
	cfg := config.Get()

	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	ctx, cancel := context.WithTimeout(ctx, cfg.TimeOutCotacaoDolar)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return models.Cotacao{}, fmt.Errorf("erro ao criar requisição: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return models.Cotacao{}, fmt.Errorf("timeout ao buscar cotação: %w", err)
		}
		return models.Cotacao{}, fmt.Errorf("erro na requisição HTTP: %w", err)
	}
	defer resp.Body.Close()

	var c models.Cotacao
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return models.Cotacao{}, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	dbCtx, cancel := context.WithTimeout(ctx, cfg.TimeOutGravaBD)
	defer cancel()
	if err := database.DB.WithContext(dbCtx).Create(&c).Error; err != nil {
		return models.Cotacao{}, fmt.Errorf("erro ao salvar no banco: %w", err)
	}

	return c, nil
}
