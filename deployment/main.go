package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Guiyu Parking!")
	})

	log.Printf("Listening and serving http on %s\n", ":8080")
	router.Run(":8080")
}
