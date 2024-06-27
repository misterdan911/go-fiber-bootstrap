package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-bootstrap/database"
	"log"
)

type User struct {
	Username string `json:"username" validate:"required" example:"danu"`
	Email    string `json:"email" validate:"required,email,email_unique" example:"dciptadi@gmail.com"`
	Password string `json:"password" validate:"required" example:"12345678"`
}

type ResponseOK struct {
	Status string `json:"status" example:"success"`
	Data   *User  `json:"data"`
}

var validate = validator.New()

func init() {
	err := validate.RegisterValidation("email_unique", emailUnique)
	if err != nil {
		log.Fatal("Failed to register custom validation 'email_unique'")
	}
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

	// Validate the user struct
	if err := validate.Struct(user); err != nil {
		// Format validation errors
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			var message string

			switch err.Tag() {
			case "email_unique":
				message = "Email has already been registered"
			default:
				//message = fmt.Sprintf("Field '%s' is invalid", err.Field())
				message = fmt.Sprintf("Field '%s' %s", err.Field(), err.Tag())
			}

			errors[err.Field()] = message
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"data":   errors,
		})
	}

	// Create
	database.DB.Create(user)

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

// Custom validation function
func emailUnique(fl validator.FieldLevel) bool {
	/*
		email := fl.Field().String()
		var user User
		if err := main.DB.Where("email = ?", email).First(&user).Error; err == nil {
			return false // email already exists
		} else if err != gorm.ErrRecordNotFound {
			return false // database error
		}
		return true // email is unique
	*/

	return true
}
