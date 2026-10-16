package response

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormValidationError(err error) any {
	var validateErr validator.ValidationErrors

	if errors.As(err, &validateErr) {
		errorMesages := make(map[string]string)
		for _, errV := range validateErr {
			field := strings.ToLower(errV.Field())
			switch errV.Tag() {
			case "required":
				errorMesages[field] = "Field " + field + " is required"
			case "email":
				errorMesages[field] = "Format email is not valid"
			case "min":
				errorMesages[field] = "Field " + field + " minimal " + errV.Param() + " character"
			default:
				errorMesages[field] = "Field " + field + " is not valid"
			}
		}

		return errorMesages
	}

	return "Payload json is not valid"
}
