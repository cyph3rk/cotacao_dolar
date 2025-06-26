// config/config.go
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	cfg      *Config
	cfgMutex sync.RWMutex
)

type Config struct {
	Timeouts struct {
		CotacaoDolar string `json:"cotacao_dolar"`
		GravaBD      string `json:"grava_bd"`
	} `json:"timeouts"`
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
	TimeOutCotacaoDolar time.Duration
	TimeOutGravaBD      time.Duration
}

func LoadConfig(path string) error {
	cfgMutex.Lock()
	defer cfgMutex.Unlock()

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var newCfg Config
	if err := json.NewDecoder(file).Decode(&newCfg); err != nil {
		return err
	}

	duration, err := time.ParseDuration(newCfg.Timeouts.CotacaoDolar)
	if err != nil {
		return fmt.Errorf("formato de duração inválido em cotacao_dolar: %v", err)
	}
	newCfg.TimeOutCotacaoDolar = duration

	duration, err = time.ParseDuration(newCfg.Timeouts.GravaBD)
	if err != nil {
		return fmt.Errorf("formato de duração inválido em grava_bd: %v", err)
	}
	newCfg.TimeOutGravaBD = duration

	cfg = &newCfg
	return nil
}

func Get() Config {
	cfgMutex.RLock()
	defer cfgMutex.RUnlock()

	if cfg == nil {
		return Config{}
	}
	return *cfg
}
