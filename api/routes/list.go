package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/s0undy/karriarum-ctf/models"
	"gorm.io/gorm"
)

func ListScore(c *fiber.Ctx, db *gorm.DB) error {
	//Grab all the data from the database and order by flags desc
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

	//Format the grabbed data into a JSON object and return it
	var parsedResults []models.Leaderboard
	if err := json.Unmarshal([]byte(jsonString), &parsedResults); err != nil {
		panic("Failed to unmarshal JSON string: " + err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(parsedResults)
}
