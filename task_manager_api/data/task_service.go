package data

import (
	"errors"
	"fmt"
	"time"

	"github.com/emrankamil/a2sv-g5-backend-track/tree/main/task_manager_api/models"
)

// Mock data for tasks
var tasks = []models.Task{}

var nextID int

func GetTasks() []models.Task{
	return tasks
}

func GetTask(id string) (models.Task, error){
	for _, task := range tasks{
		if task.ID == id{
			return task, nil
		}
	}
	
	return models.Task{}, errors.New("TASK NOT FOUND")
}
func UpdateTask(id string, updatedTask models.Task) error {
    for i, task := range tasks {
        if task.ID == id {
            tasks[i] = updatedTask
            tasks[i].ID = id
            return nil
        }
    }
	return errors.New("task not found")
}
func DeleteTask(id string) error{

	for i, task := range tasks{
		if task.ID == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}

	return errors.New("TASK NOT FOUND")
}

func PostTask(newTask models.Task){
	newTask.ID = fmt.Sprint(nextID + 1)
	newTask.DueDate = time.Now()
	nextID += 1
	tasks = append(tasks, newTask)
}