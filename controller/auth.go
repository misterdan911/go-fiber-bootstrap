package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-bootstrap/dto"
	"go-fiber-bootstrap/model"
	"go-fiber-bootstrap/service/authservice"
	"log"
)

type SignInResponseData struct {
	User model.User `json:"user"`
	Jwt  string     `json:"jwt"`
}

type SignInFailMessage struct {
	SignIn string `json:"signin"`
}

var validate = validator.New()

func init() {
	err := validate.RegisterValidation("email_unique", authservice.EmailUnique)
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
//	@Param			user	body		model.User	true	"Add user"
//	@Success		200		{object}	ResponseOK
//	@Failure		400		{object}	model.User
//	@Failure		404		{object}	model.User
//	@Failure		500		{object}	model.User
//	@Router			/signup [post]
func SignUp(c *fiber.Ctx) error {

	// Create a new User struct
	user := new(model.User)

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
	err := authservice.AddNewUser(user)
	var response interface{}

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		response = ResponseError{
			Status:  "error",
			Message: err.Error(),
		}
	} else {
		response = AppResponse{
			Status: "success",
			Data:   user,
		}
	}

	// Access the parsed user data
	// For example, you can print it or save it to the orm
	println("Name: ", user.Username)
	println("Email: ", user.Email)
	println("Password: ", user.Password)

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

// SignIn godoc
//
//	@Summary		Signing in user
//	@Description	Signing in user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		ExampleSignedInUser	true	"SignIn User"
//	@Success		200		{object}	ResponseOK
//	@Router			/signin [post]
func SignIn(c *fiber.Ctx) error {

	signedInUser := new(dto.SignedInUser)

	if err := c.BodyParser(signedInUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse SignIn JSON",
		})
	}

	isValid, userData, jwt := authservice.ValidateSignIn(signedInUser)

	var response interface{}

	/*
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	*/

	if isValid {
		resData := SignInResponseData{
			User: userData,
			Jwt:  jwt,
		}
		response = AppResponse{Status: "success", Data: resData}
	} else {
		signInFailMessage := SignInFailMessage{SignIn: "Incorrect username, email, or password"}
		response = AppResponse{
			Status: "fail",
			Data:   signInFailMessage,
		}
	}

	// supaya field2 response json nya sesuai urutan kita
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// supaya response headernya 'application/json'
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(jsonResponse)
}
