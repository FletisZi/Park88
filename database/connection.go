package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // Driver Postgres
)

func Connect(driver, source string) (*sql.DB, error) {

	db, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}

	// Configurações de Pool → muito importante!
	db.SetMaxOpenConns(25) // máximo de conexões abertas
	db.SetMaxIdleConns(25) // conexões ociosas
	db.SetConnMaxLifetime(5 * time.Minute)

	// Testa conexão
	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("📡 Banco conectado com sucesso!")
	return db, nil
}
