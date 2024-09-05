package course

import "github.com/gin-gonic/gin"

func AddRoutes(parent *gin.RouterGroup) {
	group := parent.Group("/courses")

	group.POST("", create)
	group.GET("/:id", fetchByID)
	group.PUT("/:id", update)
	group.DELETE("/:id", delete)
}
