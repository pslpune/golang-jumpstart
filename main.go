package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)
func HndlrForPing (c *gin.Context) {
	c.JSON(200, gin.H{
		"app":    "demo project",
		"author": "niranjan_Awati@psl.com",
		"date":   "December 2023",
		"msg":    "If you are able to see this, you know the telegram scraper is working fine",
	})
}
func main() {
	fmt.Println("this is from inside my program")
	log.Debug("inside my program already")

	gin.SetMode(gin.DebugMode)
	
	r := gin.Default()
	r.GET("/ping", HndlrForPing)
	r.Run(":8080")
}
