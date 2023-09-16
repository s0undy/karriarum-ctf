package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/s0undy/karriarum-ctf/models"
	"gorm.io/gorm"
)

type request struct {
	Name         string `json:"name"`
	Flags        uint64 `json:"flags"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobilenumber"`
}

type response struct {
	Status string
}

func AddScore(c *fiber.Ctx, db *gorm.DB) error {
	//Check the new request, if unable to parse JSON error out
	body := new(request)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	//Insert the new record into the database and respond with a 200 OK
	newRecord := models.Leaderboard{
		Name:         body.Name,
		Flags:        body.Flags,
		Email:        body.Email,
		MobileNumber: body.MobileNumber,
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
