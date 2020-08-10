package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hunzo/go-router/services"
)

func GetParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		text := c.Param("test")
		callFX := services.CallFunc1(text)
		c.JSON(200, gin.H{
			"info": callFX,
		})
	}
}

func GetQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		text := c.Query("text")
		callFX := services.CallFunc1(text)
		c.JSON(200, gin.H{
			"info": callFX,
		})
	}
}
