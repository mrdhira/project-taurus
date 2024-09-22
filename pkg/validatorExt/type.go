package validatorExt

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

// Password Custom Validation
// Must have at least 8 characters, 1 uppercase, 1 lowercase, 1 number, and 1 special character
// TODO: Need to update using regex
func passwordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper, hasLower, hasNumber, hasSpecial bool
	)
	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasNumber = true
		default:
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}

func New() *validator.Validate {
	if validate == nil {
		validate = validator.New()
	}

	validate.RegisterValidation("password", passwordValidation)
	return validate
}
