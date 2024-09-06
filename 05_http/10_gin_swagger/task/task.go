package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.RouterGroup) {
	group := router.Group("task")

	group.POST("/", createTask)
	group.GET("/", getAllTasks)
	group.GET("/:id", getTask)
	group.PUT("/:id", updateTask)
	group.DELETE("/:id", deleteTask)
}

// @Tags			Authentication
// @Accept			json
// @Router  		/task [post]
// @Param 			data	body	CreateRequest		true "data"
// @Success 200 {object} 		CreateResponse
func createTask(c *gin.Context) {
	// ...

	c.JSON(http.StatusOK, nil)
}

// @Tags			Authentication
// @Accept			json
// @Router  		/task/{id} [get]
// @Param 			id	path	int		true	"id"
// @Success 200 {object} GetResponse
func getTask(c *gin.Context) {
	// ...

	c.JSON(http.StatusOK, nil)
}

// @Tags			Authentication
// @Accept			json
// @Router  		/task/{id} [put]
// @Param 			id		path	int							true "id"
// @Param 			data	body	UpdateRequest		true "data"
// @Success 200
func updateTask(c *gin.Context) {
	// ...

	c.JSON(http.StatusOK, nil)
}

// @Tags			Authentication
// @Accept			json
// @Router  		/task/{id} [delete]
// @Param 			id		path	int							true "id"
// @Success 200
func deleteTask(c *gin.Context) {
	// ...

	c.JSON(http.StatusOK, nil)
}

// @Tags			Authentication
// @Accept			json
// @Router  		/task [get]
// @Param 			page	query	int			false	"page"
// @Param 			size	query	int			false	"size"
// @Param 			term	query	string		false	"term"
// @Success 200 {object} GetAllResponse
func getAllTasks(c *gin.Context) {
	// ...

	c.JSON(http.StatusOK, nil)
}
