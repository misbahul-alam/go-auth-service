// internal/validator/error_formatter.go

package validator

import "github.com/go-playground/validator/v10"

func FormatErrors(err error) map[string]string {
	errors := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {
		field := e.Field()

		switch e.Tag() {
		case "required":
			errors[field] = "This field is required"
		case "email":
			errors[field] = "Invalid email format"
		case "min":
			errors[field] = "Too short"
		case "max":
			errors[field] = "Too long"
		case "strong_password":
			errors[field] = "Password must contain 8 chars, 1 uppercase, 1 number"
		default:
			errors[field] = "Invalid value"
		}
	}

	return errors


}
