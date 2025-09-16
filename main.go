package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve the main page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Handle radio button selection
	/*
		r.POST("/popup", func(c *gin.Context) {
			selectedOption := c.PostForm("option")
			c.HTML(http.StatusOK, "popup.html", gin.H{
				"Option": selectedOption,
			})
		})*/

	r.GET("/popup", func(c *gin.Context) {
		choice := c.Query("option")
		c.HTML(http.StatusOK, "popup.html", gin.H{"Option": choice})
	})

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	r.Run(":8090") // Run server on localhost:8080
}
