package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa as rotas ultilizando as configuracoes defatlt do Gin
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Run the server
	router.Run(":8080")
}
