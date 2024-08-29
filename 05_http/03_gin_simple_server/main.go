package main

import (
	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.String(400, "hello")
}

func main() {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
