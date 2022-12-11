package errors

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
)

type AppError struct {
	err     error
	status  int
	details []string
}

func NewAppError(err error) *AppError {
	return &AppError{err: err}
}

func (ae *AppError) Build() *AppError {
	var (
		status  int
		details = make([]string, 0)
	)
	switch ae.err {
	case gorm.ErrRecordNotFound:
		status = http.StatusNotFound
		details = append(details, ae.err.Error())
		break
	default:
		status = http.StatusBadRequest
		switch ae.err.(type) {
		case validator.ValidationErrors:
			details = ae.errorTypeValidator(ae.err)
			break
		default:
			details = append(details, ae.err.Error())
			break
		}
	}

	ae.status = status
	ae.details = details

	return ae
}

func (ae *AppError) errorTypeValidator(err error) (messages []string) {
	for _, err := range err.(validator.ValidationErrors) {
		errorMsg := fmt.Sprintf(
			"Field validation for '%s' failed on the '%s' tag.",
			err.Field(),
			err.Tag(),
		)
		messages = append(messages, errorMsg)
	}
	return messages
}

func (ae *AppError) GetDetails() []string {
	if ae.details == nil {
		return []string{}
	}
	return ae.details
}

func (ae *AppError) GetStatus() int {
	return ae.status
}
