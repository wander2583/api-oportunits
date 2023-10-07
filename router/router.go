package router

import "github.com/gin-gonic/gin"

func Initialize() {
 // Inicializa as rotas ultilizando as configuracoes defatlt do Gin
	router *gin.Engine := gin.Default()
  // Definindo a rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
  router.Run(addr...: ":8080") // listen and serve on 0.0.0.0:8080
}
