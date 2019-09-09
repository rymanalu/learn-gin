package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World!"})
	})

	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.JSON(200, gin.H{"message": fmt.Sprintf("Hello, %s!", name)})
	})

	r.GET("/query", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstname", "John")
		lastName := c.DefaultQuery("lastname", "Lennon")

		c.JSON(200, gin.H{"message": fmt.Sprintf("Hello, %s %s!", firstName, lastName)})
	})

	r.Run()
}
