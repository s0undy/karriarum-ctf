package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/s0undy/karriarum-ctf/database"
	"github.com/s0undy/karriarum-ctf/models"
	"github.com/s0undy/karriarum-ctf/routes"
)

func setupRoutes(app *fiber.App) {
	app.Post("/api/v1/score", routes.AddScore)
	app.Get("/api/v1/list", routes.ListScore)
}

func main() {
	//Load config
	log.Println("Loading .env")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	//Connect to DB
	log.Println("Connecting to DB")
	db, err := database.ConnectDatabase(config)
	if err != nil {
		log.Fatal("could not load the database")
	}
	//Check if table already exsists in DB
	table := db.Migrator().HasTable(&models.Leaderboard{})
	if !table {
		log.Println("Database not initialized, creating tables")
		err = models.ImportTable(db)
		if err != nil {
			log.Fatal("Error initializing the DB")
		} else {
			log.Println("Succesfully initialized the DB")
		}
	} else {
		log.Println("Table already exsist, no need to initialized the DB.")
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
