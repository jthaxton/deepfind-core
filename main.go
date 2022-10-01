package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
    fmt.Println("Starting server")
		r := gin.Default()
		handler := Handler{}
		r.GET("/check", handler.HandleCheckVideo)
		r.Run()
}
