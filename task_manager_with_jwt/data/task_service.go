package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var tasksCollection *mongo.Collection = OpenCollection(Client, "tasks")

func GetTasks() ([]models.Task, error){
	var tasks []models.Task
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := tasksCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []models.Task{}, errors.New(err.Error())
	}

	for cur.Next(context.TODO()) {
		var task models.Task

		val := cur.Decode(&task)
		if val != nil {
			return []models.Task{}, errors.New(val.Error())
		}

		tasks = append(tasks, task)
	}

	if err := cur.Err(); err != nil {
		return []models.Task{}, errors.New(err.Error())
	}

	cur.Close(context.TODO())
	return tasks, nil
}

func GetTask(id string) (models.Task, error){
	var task models.Task
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
        return models.Task{},errors.New("INVALID ID")
    }

	filter := bson.D{{Key: "_id", Value: objID}}
	result := tasksCollection.FindOne(context.TODO(), filter).Decode(&task)
	if result != nil {
		return models.Task{}, errors.New(result.Error())
	}
	return task, result
}

func PostTask(newTask models.Task) error{
	newTask.DueDate = time.Now()
	_, err := tasksCollection.InsertOne(context.TODO(), newTask)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func UpdateTask(id string, updatedTask models.Task) error {
	objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return errors.New("INVALID ID")
    }

	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: updatedTask.Title},
			{Key: "description", Value: updatedTask.Description},
			{Key: "due_date", Value: updatedTask.DueDate},
			{Key: "status", Value: updatedTask.Status},
		}},
	}
	updateResult, result := tasksCollection.UpdateOne(context.TODO(), filter, update)
	if result != nil {
		return errors.New(result.Error())
	}
	if updateResult.MatchedCount == 0{
		return errors.New("TASK NOT FOUND")
	}
	return nil
}
func DeleteTask(id string) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return errors.New("INVALID ID")
    }

    result, err := tasksCollection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: objID}})
    if err != nil {
        return err 
    }

    if result.DeletedCount == 0 {
        return errors.New("task not found")
    }

    return nil
}

