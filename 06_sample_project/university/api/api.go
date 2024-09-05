package api

import (
	"fmt"
	"go-university/api/course"
	"go-university/api/grade"
	"go-university/api/professor"
	"go-university/api/student"
	"go-university/internal/config"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	r := gin.Default()

	v1 := r.Group("/v1")

	student.AddRoutes(v1)
	course.AddRoutes(v1)
	professor.AddRoutes(v1)
	grade.AddRoutes(v1)

	return &Router{
		r,
	}
}

func (r *Router) Serve() error {
	httpConfig := config.LoadHttpConfig()

	listenAddr := fmt.Sprintf("%s:%s", httpConfig.Url, httpConfig.Port)

	return r.Run(listenAddr)
}
