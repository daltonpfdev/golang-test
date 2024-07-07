package router

import "github.com/gin-gonic/gin"

func Inicializador() {
	// Inicializa o Router utilizando as configurações Default do gin-gonic
	r := gin.Default()

	// Definindo a rota /ping
	inicializadorDeRotas(r)

	// Iniciando o servidor na porta :8000
	r.Run(":8000")
}
