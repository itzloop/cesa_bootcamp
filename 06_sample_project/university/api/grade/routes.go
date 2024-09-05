package grade

import "github.com/gin-gonic/gin"

func AddRoutes(parent *gin.RouterGroup) {
	group := parent.Group("/grades")

	group.POST("", addNewGrade)
}
