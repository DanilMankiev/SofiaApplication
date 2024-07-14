package entity

import (
	"errors"
	"regexp"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type RegiterInput struct {
	Name      string  `json:"name"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
	Birthdate string  `json:"birthdate"`
}

func ValidateRegisterInput(input RegiterInput) error {
	birthdateRegexp := "^(?:0[1-9]|[12]\\d|3[01])([\\/.-])(?:0[1-9]|1[012])\\1(?:19|20)\\d\\d$"
	emailRegexp := "([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\\.[a-zA-Z0-9_-]+)"
	phoneRegexp := "^((8|\\+7)[\\- ]?)?(\\(?\\d{3}\\)?[\\- ]?)?[\\d\\- ]{7,10}$"

	if input.Phone == nil && input.Email == nil {
		return errors.New("email or phone required")
	}
	if input.Email != nil {
		matched, err := regexp.MatchString(emailRegexp, *input.Email)
		if err != nil {
			return err
		}
		if matched {
			return errors.New("Invalid email")
		}
	}
	if input.Phone != nil {
		matched, err := regexp.MatchString(phoneRegexp, *input.Phone)
		if err != nil {
			return err
		}
		if matched {
			return errors.New("Invalid phonw")
		}
	}
	matched, err := regexp.MatchString(birthdateRegexp, input.Birthdate)
	if err != nil {
		return err
	}
	if matched {
		return errors.New("Invalid birthdate")
	}

	return nil

}
