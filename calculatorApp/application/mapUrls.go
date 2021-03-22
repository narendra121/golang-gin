package application

import (
	"github.com/narendra121/golang-gin/calculatorApp/controller/calculatorController"
)

func MapUrls() {
	router.GET("/add", calculatorController.Addition)
	router.GET("/sub", calculatorController.Subraction)
	router.GET("/multi", calculatorController.Multiplication)

}
