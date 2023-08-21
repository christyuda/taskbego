package handlers

import (
	"context"
	"log"
	"taskbego/db"
	"taskbego/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllTasks(c *fiber.Ctx) error {
	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{"data": tasks})
}

func GetDetailTaskByID(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Task ID"})
	}

	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var detailTask models.DetailTask
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&detailTask)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Task not found"})
		}
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{"data": detailTask})
}

func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Request"})
	}

	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, task)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}

	task.ID = result.InsertedID.(primitive.ObjectID)
	return c.JSON(fiber.Map{"data": task})
}

func GetTaskByID(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Task ID"})
	}

	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var task models.Task
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Task not found"})
		}
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{"data": task})
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Task ID"})
	}

	var updatedTask models.Task
	if err := c.BodyParser(&updatedTask); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Request"})
	}

	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"completed":   updatedTask.Completed,
			"category":    updatedTask.Category,
			"priority":    updatedTask.Priority,
			"description": updatedTask.Description,
			"attachment":  updatedTask.Attachment,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}

	updatedTask.ID = objectID
	return c.JSON(fiber.Map{"data": updatedTask})
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Task ID"})
	}

	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
