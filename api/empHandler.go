package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	
)


func AddEmp() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := NewEmp()
		c.Bind(&requestBody)
		err:=AddEmpService(requestBody)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK,gin.H{"message":"employee added sucessfully"})
		}
		
	}
}


// Update Emp Update Handler
func UpdateEmp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var UpdateObj interface{}
		c.Bind(&UpdateObj)
		err:=UpdateEmpService(UpdateObj)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, gin.H{"message":"employee update sucessfully"})
		}
		
	}
}

//Search Emp Handler
func SearchEmp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var SearchObj interface{}
		c.Bind(&SearchObj)
		err,employeelist:=SearchEmpService(SearchObj)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, employeelist)
		}
		
	}
}

//List Employee
//Search Emp Handler
func ListEmp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ListObj interface{}
		c.Bind(&ListObj)
		fmt.Println("*************************************",ListObj)
		err,employeelist:=ListEmpService(ListObj)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, employeelist)
		}
		
	}
}