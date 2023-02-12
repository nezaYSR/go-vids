package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateNotNsfw(field validator.FieldLevel) bool {
	return !strings.Contains(field.Field().String(), "Nsfw")
}
