package config

// import (
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// func InitializeSQLite() (*gorm.DB, error) {
// 	// Pega o logger (se você tiver mesmo esse GetLogger implementado)
// 	logger := GetLogger("sqlite")

// 	// Abre o banco SQLite (arquivo local)
// 	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
// 	if err != nil {
// 		logger.err.Println("Erro ao conectar no SQLite:", err)
// 		return nil, err
// 	}

// 	logger.info.Println("SQLite conectado com sucesso!")
// 	return db, nil
// }
