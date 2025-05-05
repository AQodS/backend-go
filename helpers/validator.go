package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func TranslateErrorMessage(err error) map[string]string {
	// create map to store error messages
	errorsMap := make(map[string]string)

	// handle validate from validator
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			field := fieldError.Field() // save field name
			switch fieldError.Tag() {
			case "required":
				errorsMap[field] = fmt.Sprintf("%s is required", field)
			case "email":
				errorsMap[field] = "invalid email format"
			case "unique":
				errorsMap[field] = fmt.Sprintf("%s already exists", field)
			case "min":
				errorsMap[field] = fmt.Sprintf("%s must be at least %s characters", field, fieldError.Param())
			case "max":
				errorsMap[field] = fmt.Sprintf("%s must be at most %s characters", field, fieldError.Param())
			case "numeric":
				errorsMap[field] = fmt.Sprintf("%s must be numeric", field)
			default:
				errorsMap[field] = "invalid value"
			}
		}
	}

	// handle error from gorm for duplicate entry
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "username") {
				errorsMap["Username"] = "Username already exists"
			}
			if strings.Contains(err.Error(), "email") {
				errorsMap["Email"] = "Email already exists"
			}
		} else if err == gorm.ErrRecordNotFound {
			// if data isn't found in database
			errorsMap["Error"] = "Record not found"
		}
	}

	return errorsMap // return error messages
}

func IsDuplicateEntryError(err error) bool {
	// check if error is duplicate entry error
	return err != nil && strings.Contains(err.Error(), "Duplicate entry")
}
