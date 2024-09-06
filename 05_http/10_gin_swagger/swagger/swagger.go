package swagger

import (
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func AddRoutes(router *gin.RouterGroup) {
	router.GET("/swagger/*any", swagger.WrapHandler(files.Handler))
}
