package main

import (
	"fmt"
	"net/http"
	"EmployeeAssisgnment/api"
	"github.com/gin-gonic/gin"
	helper "EmployeeAssisgnment/api/helpers"
)

func main() {
	fmt.Println("Gin-Gonic Server")
	 helper.InitDB()
	startServer()
}

func startServer() {
	router := gin.Default()
	router.GET("/", checkStatus())
	api.Init(router)
	s := &http.Server{
		Addr:    ":4700",
		Handler: router,
	}
	s.ListenAndServe()
}

func checkStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is running successfully !!!!!")
	}
}
