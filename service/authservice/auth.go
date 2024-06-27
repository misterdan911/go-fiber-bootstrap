package authservice

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"go-fiber-bootstrap/model"
	"go-fiber-bootstrap/orm"
	"golang.org/x/crypto/bcrypt"
)

func AddNewUser(user *model.User) error {

	user.Password, _ = HashPassword(user.Password)
	result := orm.DB.Create(user)

	// Check for errors
	if result.Error != nil {
		return errors.New("Failed creating new user, " + result.Error.Error())
	} else {
		return nil
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Custom validation function
func EmailUnique(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	var user model.User

	result := orm.DB.Find(&user, "email = ?", email)

	if result.RowsAffected == 0 {
		return true
	} else {
		return false
	}
}
