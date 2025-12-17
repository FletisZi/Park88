package main

import (
	"github.com/fletiszi/goteste/config"
	router "github.com/fletiszi/goteste/router"
)

// var (
// 	logger *config.Logger
// )

func main() {

	// logger = config.GetLogger("main")
	// err := config.Init()

	// if err != nil {
	// 	logger.Errf("Ocorreu erro ao iniciar config.init %v", err)
	// 	return
	// }

	// cfg := config.LoadConfig()

	// conn, err := database.Connect(cfg.DBDriver, cfg.DBSource)
	// if err != nil {
	// 	logger.Errf("Erro ao conectar no banco:%v", err)
	// }

	// defer conn.Close()

	config.LoadEnv()

	// Inicializa banco de dados
	config.InitDatabase()

	router.Initialize()

}
