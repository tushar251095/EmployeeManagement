package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	service "EmployeeAssisgnment/api/services"
)


func AddEmp() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := service.NewEmp()
		c.Bind(&requestBody)
		err:=service.AddEmpService(requestBody)
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
		err:=service.UpdateEmpService(UpdateObj)
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
		err,employeelist:=service.SearchEmpService(SearchObj)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, employeelist)
		}
		
	}
}

//List Employee
func ListEmp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ListObj interface{}
		c.Bind(&ListObj)
		err,employeelist:=service.ListEmpService(ListObj)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, employeelist)
		}
		
	}
}

//Delete Employee
func DeleteEmp() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody:=service.NewDeleteData()
		c.Bind(&requestBody)
		err,msg:=service.DeleteEmpService(requestBody)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, gin.H{"message":msg})
		}
		
	}
}

//Restore Employee
func RestoreEmp() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody:=service.NewRestoreData()
		c.Bind(&requestBody)
		err,msg:=service.RestoreEmpService(requestBody)
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, gin.H{"message":msg})
		}
		
	}
}

//VIEW DELETED EMPLOYEE
func ViewDeletedEmp() gin.HandlerFunc {
	return func(c *gin.Context) {
		err,employeelist:=service.ViewDeletedEmpService()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
		}else{
			c.JSON(http.StatusOK, employeelist)
		}
		
	}
}