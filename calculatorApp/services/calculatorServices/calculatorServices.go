package calculatorServices

import "github.com/narendra121/golang-gin/calculatorApp/domain/calculator"

func Add(add calculator.Calculator) (int64) {
		return add.Number1+add.Number2
}

func Sub(add calculator.Calculator) (int64) {
	return add.Number1-add.Number2
}

func Multi(add calculator.Calculator) (int64) {
	return add.Number1*add.Number2
}