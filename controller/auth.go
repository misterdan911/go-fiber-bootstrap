package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
)

type User struct {
	Username string `json:"username" example:"danu"`
	Email    string `json:"email" example:"dciptadi@gmail.com"`
	Password string `json:"password" example:"12345678"`
}

type ResponseOK struct {
	Status string `json:"status" example:"success"`
	Data   *User  `json:"data"`
}

// SignUp godoc
//
//	@Summary		Signing up new user
//	@Description	Signing up new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		User	true	"Add user"
//	@Success		200		{object}	ResponseOK
//	@Failure		400		{object}	User
//	@Failure		404		{object}	User
//	@Failure		500		{object}	User
//	@Router			/signup [post]
func SignUp(c *fiber.Ctx) error {

	// Create a new User struct
	user := new(User)

	// Parse the JSON request body into the user struct
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	// Access the parsed user data
	// For example, you can print it or save it to the database
	println("Name: ", user.Username)
	println("Email: ", user.Email)
	println("Password: ", user.Password)

	response := ResponseOK{
		Status: "success",
		Data:   user,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Send a response
	/*
		return c.JSON(fiber.Map{
			"status": "success",
			"data":   user,
		})
	*/

	//return c.JSON(jsonResponse)

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(jsonResponse)

}