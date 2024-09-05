package grade

import (
	"go-university/internal/app/grade"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addNewGrade(c *gin.Context) {
	var payload grade.Grade
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	id, err := grade.AddNewGrade(ctx, &payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, id)
}
