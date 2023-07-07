package main

import (
	"context"
	"github.com/AhmedSaIah/fiber/controllers"
	"github.com/AhmedSaIah/fiber/database"
	"github.com/AhmedSaIah/fiber/repository"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("MONGO_URL")
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(database.NewConnection())
	userController := controllers.NewUserController(userRepo)

	app.Use(logger.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Hello fiber"})
	})
	app.Post("/signup", userController.SignUp)
	//app.Post("/login", userController.Login)
	log.Fatal(app.Listen(":3000"))
}
