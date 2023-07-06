package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/AhmedSaIah/fiber/controllers/user"
	"github.com/AhmedSaIah/fiber/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("MONGO_URL")
	db, err := database.ConnDB(url)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(nil)

	app := fiber.New(fiber.Config{Immutable: true})
	userRepo := database.NewUserRepository(db)
	signUpController := user.NewUserController(userRepo)

	app.Get("/")
	app.Post("/signup", signUpController.SignUp)
	app.Post("/login", signUpController.Login)
	err = app.Listen(":3000")
	if err != nil {
		fmt.Println("listening on port 3000 failed %w", err)
	}
}
