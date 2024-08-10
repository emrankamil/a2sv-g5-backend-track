package router

import (
	"net/http"
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func RoutingHandler(r *gin.Engine) {
	
	r.GET("/ping", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//tasks
	r.GET("/tasks", controllers.GetTasks)
	r.GET("/tasks/:id", controllers.GetTaskById)
	r.PUT("/tasks/:id", middleware.AuthMiddleware(), middleware.AuthRole("ADMIN"), controllers.UpdateTask)
	r.DELETE("/tasks/:id", middleware.AuthMiddleware(), middleware.AuthRole("ADMIN"), controllers.DeleteTask)
	r.POST("/tasks", middleware.AuthMiddleware(), middleware.AuthRole("ADMIN"), controllers.PostTask)

	//users
	r.POST("/register", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.PUT("/promote/:id", middleware.AuthMiddleware(), middleware.AuthRole("ADMIN"), controllers.PromoteUser)
	r.GET("/users", middleware.AuthMiddleware(), middleware.AuthRole("ADMIN"),controllers.GetUsers)
}