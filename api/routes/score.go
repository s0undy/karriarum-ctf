package routes

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/s0undy/karriarum-ctf/database"
	"github.com/s0undy/karriarum-ctf/models"
)

type request struct {
	Name  string `json:"name"`
	Flags uint64 `json:"flags"`
}

type response struct {
	Status string
}

func AddScore(c *fiber.Ctx) error {
	//Check the new request, if unable to parse JSON error out
	body := new(request)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	//Connect to DB
	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := database.ConnectDatabase(config)
	if err != nil {
		log.Fatal("could not load the database")
	}
	//Insert the new record into the database and respond with a 200 OK
	newRecord := models.Leaderboard{
		Name:  body.Name,
		Flags: body.Flags,
	}
	db.Create(&newRecord)
	if db.Error != nil {
		panic("Failed to create a new record " + db.Error.Error())
	}
	resp := response{
		Status: "OK",
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
