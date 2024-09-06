// git clone https://github.com/kkoutsilis/todo-go-htmx

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"todo-go-htmx/database"
	"todo-go-htmx/handlers"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

const fileName = "sqlite.db"

func main() {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}

	repository := database.NewSQLiteRepository(db)

	if err := repository.Migrate(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migrated")

	router := gin.Default()
	router.HTMLRender = &TemplRender{}

	homeHtmlHandler := handlers.HomeHtmlHandler{Repository: repository}
	todoHtmlHandler := handlers.TodoHtmlHandler{Repository: repository}

	router.GET("/", homeHtmlHandler.GetHome)

	router.GET("/todo", todoHtmlHandler.GetAll)
	router.POST("/todo", todoHtmlHandler.Create)
	router.PATCH("/todo/:id", todoHtmlHandler.Update)
	router.DELETE("/todo/:id", todoHtmlHandler.Delete)

	router.Run(":8080")
}

type TemplRender struct {
	Code int
	Data templ.Component
}

func (t TemplRender) Render(w http.ResponseWriter) error {
	t.WriteContentType(w)
	w.WriteHeader(t.Code)
	if t.Data != nil {
		return t.Data.Render(context.Background(), w)
	}
	return nil
}

func (t TemplRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func (t *TemplRender) Instance(name string, data interface{}) render.Render {
	if templData, ok := data.(templ.Component); ok {
		return &TemplRender{
			Code: http.StatusOK,
			Data: templData,
		}
	}
	return nil
}
