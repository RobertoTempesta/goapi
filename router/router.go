package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicia o Router utilizando as configurações Default do gin
	router := gin.Default()
	//Define uma Rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Estamos rodando a nossa api
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
