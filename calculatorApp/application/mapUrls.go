package application

import (
	"github.com/narendra121/golang-gin/calculatorApp/controller/calculatorController"
)

func MapUrls() {
	router.POST("/add", calculatorController.Addition)
	router.POST("/sub", calculatorController.Subraction)
	router.POST("/multi", calculatorController.Multiplication)

}
