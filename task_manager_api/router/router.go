package router

import (
	"net/http"
	"github.com/emrankamil/a2sv-g5-backend-track/tree/main/task_manager_api/controllers"
	"github.com/gin-gonic/gin"
)

func RoutingHandler(r *gin.Engine) {
	
	r.GET("/ping", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/tasks", controllers.GetTasks)
	r.GET("/tasks/:id", controllers.GetTaskById)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
	r.POST("/tasks", controllers.PostTask)
}