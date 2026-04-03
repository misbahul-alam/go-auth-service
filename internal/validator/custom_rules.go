package validator

import "github.com/go-playground/validator/v10"

func registerCustomRules() {
	err := Validate.RegisterValidation("strong_password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()

		if len(password) < 8 {
			return false
		}

		hasNumber := false
		hasUpper := false

		for _, c := range password {
			if c >= '0' && c <= '9' {
				hasNumber = true
			}
			if c >= 'A' && c <= 'Z' {
				hasUpper = true
			}
		}

		return hasNumber && hasUpper
	})
	if err != nil {
		return
	}
}
