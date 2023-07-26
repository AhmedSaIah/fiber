package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/AhmedSaIah/fiber/config"
	"github.com/AhmedSaIah/fiber/controllers"
	"github.com/AhmedSaIah/fiber/routes"
	"github.com/AhmedSaIah/fiber/services"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoClient *mongo.Client
	redisClient *redis.Client

	userService         services.UserService
	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	authCollection      *mongo.Collection
	authService         services.AuthService
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController
)

func init() {
	//cnf, err := config.LoadConfig(".")
	//if err != nil {
	//	log.Fatal("Could not load environment variables", err)
	//}

	ctx = context.TODO()

	// Connect to MongoDB
	mongoConn := options.Client().ApplyURI("mongodb+srv://admin:GGv0Oru3qvZ15CYJ@fiber.w9k8k3m.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(ctx, mongoConn)
	if err != nil {
		panic(err)
	}

	//if err := client.Ping(ctx, readpref.Primary()); err != nil {
	//	panic(err)
	//}

	fmt.Println("MongoDB successfully connected...")

	// Connect to Redis
	//redisClient := redis.NewClient(&redis.Options{
	//	Addr: cnf.RedisUri,
	//})
	//
	//if _, err := redisClient.Ping(ctx).Result(); err != nil {
	//	panic(err)
	//}
	//
	//err = redisClient.Set(ctx, "test", "Welcome to Golang with Redis and MongoDB", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("Redis client connected successfully...")

	// Collections
	authCollection = client.Database("fiber").Collection("users")
	userService = services.NewUserServiceImpl(authCollection, ctx)
	authService = services.NewAuthService(authCollection, ctx)
	AuthController = controllers.NewAuthController(authService, userService)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(userService)
	UserRouteController = routes.NewRouteUserController(UserController)

	// Assign the Fiber app to the global variable if necessary
	// app := fiber.New(fiber.Config{
	// 	ErrorHandler: customErrorHandler,
	// })

	server = gin.Default()
}
func main() {
	cnf, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load cnf", err)
	}

	defer mongoClient.Disconnect(ctx)

	//value, err := redisClient.Get(ctx, "test").Result()

	//if err == redis.Nil {
	//	fmt.Println("key: test does not exist")
	//} else if err != nil {
	//	panic(err)
	//}

	//TODO: Install cors gor gin-gonic and almost done

	corsConfig := cors.Config{
		AllowOrigins:     "http://localhost:8000",
		AllowCredentials: true,
	}
	server.Use(cors.New(corsConfig))

	//app.Get("/", func(ctx *gin.Context) {
	//	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "/"})
	//})

	router := app.Group("/auth")
	router.Get("/check", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "works"})
	})

	AuthRouteController.AuthRoute(router, userService)
	UserRouteController.UserRoute(router, userService)

	err = app.Listen(":" + cnf.Port)
	if err != nil {
		log.Fatal(err)
	}

}
