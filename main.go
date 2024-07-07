package main

import (
	"github.com/daltonpfdev/golang-test/config"
	"github.com/daltonpfdev/golang-test/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Iniciando as Configurações
	err := config.InitConfig()
	if err != nil {
		logger.Errf("Erro na Inicialização das Configurações: %v", err)
		return
	}

	// Iniciando o Router
	router.Inicializador()
}
