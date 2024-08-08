package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"
	"github.com/gin-gonic/gin"
)


func GetTasks(c *gin.Context){
	tasks, err := data.GetTasks()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
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
	result := data.PostTask(newTask)
	if result != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":result.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task Posted"})
}


func UpdateTask(c *gin.Context){
	id := c.Param("id")
	var updatedTask models.Task

	result := c.ShouldBindJSON(&updatedTask)
	if result != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":result.Error()})
		return
	}

	err := data.UpdateTask(id, updatedTask)
	if err != nil{
		if err.Error() == "TASK NOT FOUND"{
			c.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "TASK UPDATED"})
}

func DeleteTask(c *gin.Context){
	id := c.Param("id")

	result := data.DeleteTask(id)
	if result != nil{
		if result.Error() == "INVALID ID"{
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": result.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

