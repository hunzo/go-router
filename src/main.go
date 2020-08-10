package main

import (
	"fmt"

	"github.com/hunzo/go-router/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("test")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api1")
	{
		v1.Use(midleWare())
		v1.GET("/getparam/:test", handler.GetParams())
	}

	v2 := r.Group("/api2")
	{
		v2.GET("/getquery", handler.GetQuery())
	}

	r.Run()

}

func midleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Hi Middleware")
		fmt.Printf("%v\n", c.Request.Header)
		c.Next()

	}
}
