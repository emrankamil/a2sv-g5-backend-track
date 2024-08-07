package controllers

import (
	"net/http"
	"github.com/emrankamil/a2sv-g5-backend-track/tree/main/task_manager_api/models"
	"github.com/emrankamil/a2sv-g5-backend-track/tree/main/task_manager_api/data"
	"github.com/gin-gonic/gin"
)


func GetTasks(c *gin.Context){
	tasks := data.GetTasks()
	c.JSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context){
	id := c.Param("id")

	task, err := data.GetTask(id)
	if err == nil{
		c.JSON(http.StatusOK, task)
			return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
}

func PostTask(c *gin.Context){
	var newTask models.Task

	err := c.ShouldBindJSON(&newTask)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	data.PostTask(newTask)
	c.JSON(http.StatusOK, gin.H{"message": "Task Posted"})
}


func UpdateTask(c *gin.Context){
	id := c.Param("id")
	var updatedTask models.Task

	err := c.ShouldBindJSON(&updatedTask)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	data.UpdateTask(id, updatedTask)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "TASK UPDATED"})
}

func DeleteTask(c *gin.Context){
	id := c.Param("id")
	err := data.DeleteTask(id)

	if err == nil{
		c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
}

