package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/s0undy/karriarum-ctf/database"
	"github.com/s0undy/karriarum-ctf/models"
	"github.com/s0undy/karriarum-ctf/routes"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App, db *gorm.DB) {
	app.Post("/api/v1/score", func(c *fiber.Ctx) error {
		return routes.AddScore(c, db)
	})
	app.Get("/api/v1/list", func(c *fiber.Ctx) error {
		return routes.ListScore(c, db)
	})
}

func main() {
	//Check that all environment variables are set
	//log.Println("Loading .env")
	//err := godotenv.Load("../.env")
	//if err != nil {
	//	log.Fatal(err)
	//}
	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	if config.Host == "" {
		log.Fatal("DB_HOST environment variable is not set")
	}
	if config.Port == "" {
		log.Fatal("DB_PORT environment variable is not set")
	}
	if config.Password == "" {
		log.Fatal("DB_PASSWORD environment variable is not set")
	}
	if config.User == "" {
		log.Fatal("DB_USER environment variable is not set")
	}
	if config.SSLMode == "" {
		log.Fatal("DB_SSLMode environment variable is not set")
	}
	if config.DBName == "" {
		log.Fatal("DB_DBName environment variable is not set")
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

	setupRoutes(app, db)

	app_port := os.Getenv("APP_PORT")
	if app_port == "" {
		log.Fatal("APP_PORT environment variable is not set")
	}
	log.Fatal((app.Listen(app_port)))
	//log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
