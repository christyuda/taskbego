package main

import (
	"log"
	"os"
	"taskbego/db"
	"taskbego/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Mengambil nilai environment variable MONGODB_URI
	mongoURI := os.Getenv("mongodb+srv://chrisyuda:yudacihuy123@clusterdb.xld4k4m.mongodb.net/")

	if mongoURI == "" {
		log.Fatal("MONGODB_URI environment variable is not set")
	}

	err := db.ConnectMongoDB(mongoURI, "dbgolang", "tasks")
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	app.Get("/tasks", handlers.GetAllTasks)
	app.Post("/tasks", handlers.CreateTask)
	app.Get("/tasks/:id", handlers.GetTaskByID)
	app.Put("/tasks/:id", handlers.UpdateTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)
	app.Get("/tasks/:id/detail", handlers.GetDetailTaskByID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port
	}

	log.Fatal(app.Listen(":" + port))
}
