package main

import (
	"fmt"
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
    fmt.Println("Starting server")
		r := gin.Default()
		engine := Engine{}
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		r.GET("/find", engine.HandleCheckVideo)
		r.Run()
}