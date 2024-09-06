package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// $ curl "localhost:8080/test"

var users = gin.Accounts{
	"admin": "admin",
	"test":  "123",
}

func main() {
	r := gin.Default()
	r.Use(gin.BasicAuth(users))
	r.Use(Logger())
	r.Use(Timeout())

	r.GET("/test", func(c *gin.Context) {
		userID := c.GetString("userID")

		time.Sleep(1 * time.Second)

		log.Println(userID)
	})

	r.Run(":8080")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set example variable
		c.Set("userID", "1")

		t := time.Now()
		// // before request
		c.Next()

		fmt.Printf("duration: %vus\n", time.Since(t).Microseconds())

		// // access the status we are sending
		status := c.Writer.Status()
		log.Println("Status Code:", status)
	}
}

func Timeout() gin.HandlerFunc {
	return func(c *gin.Context) {
		finish := make(chan struct{})

		go func() {
			c.Next()
			finish <- struct{}{}
		}()

		select {
		case <-time.After(2 * time.Second):
			c.JSON(http.StatusGatewayTimeout, "timeout")
			c.Abort()
		case <-finish:
		}
	}
}

// var users = gin.Accounts{
// 	"admin": "admin",
// 	"test":  "123",
// }

// r.Use(gin.BasicAuth(users))

// $ curl "localhost:8080/test" -u admin:admin
