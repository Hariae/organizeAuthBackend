package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

// credential represents data on user credentials
type credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Service  string `json:"service"`
}

var credentials = []credential{
	{Username: "aehari2010@gmail.com", Password: "password@123", Service: "netflix"},
	{Username: "aehari2010@gmail.com", Password: "password@123", Service: "icici"},
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, nil)
	})
	router.GET("/credentials", getCredentials)
	router.Run(":" + port)
}

func getCredentials(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, credentials)
}
