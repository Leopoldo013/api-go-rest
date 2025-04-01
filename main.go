package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.GET("/events", getEvents)
	server.Run(":8080")
}

func getEvents(c *gin.Context) {
	events := []string{"Event 1", "Event 2", "Event 3"}
	c.JSON(200, events)
}
