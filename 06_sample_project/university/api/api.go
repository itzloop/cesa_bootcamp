package api

import (
	"fmt"
	"go-university/internal/config"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	r := gin.Default()

	return &Router{
		r,
	}
}

func (r *Router) Serve() error {
	httpConfig := config.LoadHttpConfig()

	listenAddr := fmt.Sprintf("%s:%s", httpConfig.Url, httpConfig.Port)

	return r.Run(listenAddr)
}
