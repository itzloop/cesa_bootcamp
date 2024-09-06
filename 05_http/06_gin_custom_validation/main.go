package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// $ curl "localhost:8080/bookable?check_in=2118-04-16&check_out=2118-04-17"

// $ curl "localhost:8080/bookable?check_in=2118-03-10&check_out=2118-03-9"

// $ curl "localhost:8080/bookable?check_in=2018-03-10&check_out=2018-03-11"

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,bookabledate,gtfield=CheckIn" time_format:"2006-01-02"`
	Count    int64     `form:"count" binding:"required,naturalNumber"`
}

func main() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
		v.RegisterValidation("naturalNumber", naturalNumber)
	}

	route := gin.Default()

	route.GET("/bookable", getBookable)

	route.Run(":8080")
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// today := time.Now()

	// if today.After(b.CheckIn) {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "a"})
	// 	return
	// }

	// if today.After(b.CheckOut) {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "a"})
	// 	return
	// }

	// if b.CheckOut.After(b.CheckIn) {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "a"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}

	return true
}

var naturalNumber validator.Func = func(fl validator.FieldLevel) bool {
	value := fl.Field().Int()

	return value > 0
}

// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
// 	v.RegisterValidation("bookabledate", bookableDate)
// }
