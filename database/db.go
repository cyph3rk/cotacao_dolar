package database

import (
	"log"

	"github.com/cyph3rk/cotacao_dolar/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	// github.com/mattn/go-sqlite3
	// You can also use file::memory:?cache=shared instead of a path to a file.
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Cotacao{}) // gera a tabela aluno automaticamente se ela nao existe
}
