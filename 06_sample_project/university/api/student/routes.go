package student

import "github.com/gin-gonic/gin"

func AddRoutes(parent *gin.RouterGroup) {
	group := parent.Group("/students")

	group.POST("", create)
	group.GET("/:id", fetchByID)
	group.PUT("/:id", update)
	group.DELETE("/:id", delete)
}
