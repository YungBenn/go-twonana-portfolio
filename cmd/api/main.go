package main

import (
	"log"

	"github.com/YungBenn/go-twonana-portfolio/config"
	"github.com/YungBenn/go-twonana-portfolio/internal/api/routes"
	"github.com/YungBenn/go-twonana-portfolio/internal/mongodb"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	mongoStore "github.com/gofiber/storage/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	env, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err)
	}

	db, err := connectMongo(env)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %s", err)
	}
	defer func(db *mongo.Database) {
		err := mongodb.CloseMongoDB(db)
		if err != nil {
			log.Fatalf("Error closing MongoDB: %s", err)
		}
	}(db)

	app := setupApp()

	store := setupSessionStore(env)

	routes.SetupNFTRoutes(app, db, env, store)
	routes.SetupAuthRoutes(app, db, store)

	err = app.Listen(":" + env.PORT)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

	log.Printf("Server is running on port %s", env.PORT)
}

func connectMongo(env config.EnvVars) (*mongo.Database, error) {
	db, err := mongodb.ConnectDB(env.MONGODBURI, env.MONGODBNAME)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func setupApp() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New())
	app.Use(logger.New())

	app.Static("/docs", "./docs/swagger")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("OK!")
	})

	return app
}

func setupSessionStore(env config.EnvVars) *session.Store {
	sessionStore := mongoStore.New(mongoStore.Config{
		ConnectionURI: env.MONGODBURI,
		Database:      env.MONGODBNAME,
		Collection:    "sessions",
		Reset:         false,
	})

	store := session.New(session.Config{
		Storage: sessionStore,
	})

	return store
}
