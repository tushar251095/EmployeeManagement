package route

import (
	"github.com/gin-gonic/gin"
	handler "EmployeeAssisgnment/api/handler"
)

func Init(r *gin.Engine) {

	r.POST("/add", handler.AddEmp())
	r.PUT("/update", handler.UpdateEmp())
	r.POST("/search", handler.SearchEmp())
	r.POST("/list", handler.ListEmp())
	r.POST("/delete", handler.DeleteEmp())
	r.POST("/restore",handler.RestoreEmp())
	r.GET("/viewdeleted",handler.ViewDeletedEmp())
}
