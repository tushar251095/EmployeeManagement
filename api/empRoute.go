package api

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	r.POST("/add", AddEmp())
	r.PUT("/update", UpdateEmp())
	r.POST("/search", SearchEmp())
	r.POST("/list", ListEmp())
}
