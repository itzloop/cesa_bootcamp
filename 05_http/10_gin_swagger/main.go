// swag init -o 005-gin-swagger/api -d 005-gin-swagger --parseDependency --parseVendor
// http://localhost:8080/api/swagger/index.html
package main

import (
	"fmt"
	"playground/005-gin-swagger/swagger"
	"playground/005-gin-swagger/task"

	"github.com/gin-gonic/gin"

	_ "playground/005-gin-swagger/api"
)

// @title           API Docs
// @version         1.0
// @description     Golang Swagger Architectures endpoints

// @BasePath  /api

// @securityDefinitions.basic  BasicAuth
func main() {
	fmt.Println("client starting...")

	engine := gin.Default()

	group := engine.Group("api")
	swagger.AddRoutes(group)
	task.AddRoutes(group)

	engine.Run(":8080")
}
