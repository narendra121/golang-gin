package calculatorController

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/narendra121/golang-gin/calculatorApp/domain/calculator"
	"github.com/narendra121/golang-gin/calculatorApp/services/calculatorServices"
)
func Addition(c *gin.Context){
	var calculator calculator.Calculator
	if err:=c.ShouldBindJSON(&calculator);err!=nil{
		fmt.Println(err)
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid json body..."})
		return
	}
	sum:=calculatorServices.Add(calculator)
	c.JSON(http.StatusOK,gin.H{"Sum Of Given Two Number":sum})
}

func Subraction(c *gin.Context){
	var calculator calculator.Calculator
	if err:=c.ShouldBindJSON(&calculator);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid json body..."})
	}
	sum:=calculatorServices.Sub(calculator)
	c.JSON(http.StatusOK,gin.H{"Subraction Of Given Two Number":sum})
}

func Multiplication(c *gin.Context){
	var calculator calculator.Calculator
	if err:=c.ShouldBindJSON(&calculator);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid json body..."})
	}
	sum:=calculatorServices.Multi(calculator)
	c.JSON(http.StatusOK,gin.H{"multiplication Of Given Two Number":sum})
}

