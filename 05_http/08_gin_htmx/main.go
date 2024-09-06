package main

import (
	"fmt"
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jaswdr/faker/v2"
)

type DommyLOG struct {
	ID      string
	Time    string
	Content string
}

var logs []DommyLOG

func init() {
	logs = make([]DommyLOG, 0, 100)
}

var tmpl = template.Must(template.ParseFiles("layout.html"))

func main() {
	go func() {
		for {
			LOG()
			time.Sleep(1 * time.Second)
		}
	}()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		tmpl.Execute(ctx.Writer, gin.H{
			"Logs": logs,
		})
	})

	r.Run(":8080")
}

func LOG() {
	l := DommyLOG{
		ID:      uuid.NewString(),
		Time:    time.Now().Format(time.TimeOnly),
		Content: fmt.Sprintf("request Ipv4: %s, from: %s", faker.New().Internet().Ipv4(), faker.New().Address().Country()),
	}

	logs = append([]DommyLOG{l}, logs...)

	if len(logs) == 100 {
		logs = logs[:100]
	}
}
