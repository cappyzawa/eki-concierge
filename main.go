package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupRouter - ルーターの初期化を行う
func SetupRouter() *gin.Engine {
	router := gin.New()
	v1 := router.Group("/v1/")
	{
		v1.GET("search", Search)
	}
	return router
}

// Search - 検索を行う
func Search(c *gin.Context) {
	Response(c, nil)
}

// Response - Responseを返す
func Response(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"result": "hello",
	})
}

func main() {
	router := SetupRouter()
	log.Printf("Start Go HTTP Server")

	port := os.Getenv("PORT")
	router.Run(":" + port)
}
