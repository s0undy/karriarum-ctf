package routes

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/s0undy/karriarum-ctf/database"
	"github.com/s0undy/karriarum-ctf/models"
)

func ListScore(c *fiber.Ctx) error {
	//Connect to DB
	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	log.Println("Connecting to DB")
	db, err := database.ConnectDatabase(config)
	if err != nil {
		log.Fatal("could not load the database")
	}
	var result []models.Leaderboard

	db.Order("Flags desc").Find(&result)
	if db.Error != nil {
		panic("Failed to query the database: " + db.Error.Error())
	}

	jsonResults, err := json.Marshal(result)
	if err != nil {
		panic("Failed to marshal results to JSON: " + err.Error())
	}

	jsonString := string(jsonResults)

	var parsedResults []models.Leaderboard
	if err := json.Unmarshal([]byte(jsonString), &parsedResults); err != nil {
		panic("Failed to unmarshal JSON string: " + err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(parsedResults)
}
