package main

import (
	"log"
	"taskbego/db"
	"taskbego/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	err := db.ConnectMongoDB("mongodb+srv://chrisyuda:yudacihuy123@clusterdb.xld4k4m.mongodb.net/", "dbgolang", "tasks")
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	app.Get("/tasks", handlers.GetAllTasks)
	app.Post("/tasks", handlers.CreateTask)
	app.Get("/tasks/:id", handlers.GetTaskByID)
	app.Put("/tasks/:id", handlers.UpdateTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)
	app.Get("/tasks/:id/detail", handlers.GetDetailTaskByID)

	log.Fatal(app.Listen(":3000"))
}
