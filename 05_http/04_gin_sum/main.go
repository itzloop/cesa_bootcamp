package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

// api for getting two number and returning the sum of them

// example:

// Request:
// {
//	first_number: 10,
// 	second_number: 11
//	}

// Response:
// {
// sum:21
//	}

type RequestQuery struct {
	FirstNumber  int `form:"first_number"`
	SecondNumber int `form:"second_number"`
}

type Request struct {
	FirstNumber  int `json:"first_number"`
	SecondNumber int `json:"second_number"`
}

type Response struct {
	Sum int `json:"sum"`
}

func Sum(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// unmarshall the req
	var req Request
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(req)
	sum := req.FirstNumber + req.SecondNumber

	// marshal the response
	resp := Response{Sum: sum}
	bodyBytes, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.String(200, string(bodyBytes))
}

func SumUsingGinUtility(c *gin.Context) {
	var req Request
	err := c.BindJSON(&req)
	if err != nil {
		fmt.Println(err)
		return
	}
	sum := req.FirstNumber + req.SecondNumber
	c.JSON(200, Response{Sum: sum})
}

func SumQuery(c *gin.Context) {
	var req RequestQuery
	err := c.BindQuery(&req)
	if err != nil {
		fmt.Println(err)
		return
	}
	if req.FirstNumber == 0 && req.SecondNumber == 0 {
		c.String(400, "error")
		return
	}
	sum := req.FirstNumber + req.SecondNumber
	c.JSON(200, Response{Sum: sum})
}

func main() {
	r := gin.Default()
	r.POST("/sum", Sum)
	r.POST("/sum2", SumUsingGinUtility)
	r.POST("/sumQuery", SumQuery)
	r.Run(":8080")
}
