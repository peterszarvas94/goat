package validation

import (
	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	Validator = validator.New()
}

func ValidateStruct(s any) error {
	return Validator.Struct(s)
}

// func ValidateVar(field any, tag string) error {
// 	return Validator.Var(field, tag)
// }

// BuildValidationMessages converts validation errors to user-friendly messages
func BuildValidationMessages(err error) []string {
	var messages []string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			switch fieldErr.Tag() {
			case "required":
				messages = append(messages, fieldErr.Field()+" is required")
			case "min":
				messages = append(messages, fieldErr.Field()+" must be at least "+fieldErr.Param()+" characters")
			case "max":
				messages = append(messages, fieldErr.Field()+" must be at most "+fieldErr.Param()+" characters")
			case "email":
				messages = append(messages, "Email must be a valid email address")
			case "alphanum":
				messages = append(messages, fieldErr.Field()+" must contain only letters and numbers")
			default:
				messages = append(messages, fieldErr.Field()+" is invalid")
			}
		}
	} else {
		messages = append(messages, err.Error())
	}

	return messages
}

// Custom validation middleware
// func ValidationMiddleware(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		r.Body = http.MaxBytesReader(w, r.Body, 1048576) // 1MB limit
// 		next(w, r)
// 	}
// }
