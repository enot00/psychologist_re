package validators

import (
	"github.com/go-playground/validator"
	"github.com/test_server/internal/infra/http/resources"
	"log"
)

type ChatValidator struct {
	validator *validator.Validate
}

func NewChatValidator() *ChatValidator {
	return &ChatValidator{
		validator: validator.New(),
	}
}

func (v *ChatValidator) Validate(dto *resources.ChatDto) error {
	err := v.validator.Struct(dto)
	if err != nil {
		log.Printf("validation error: %v", err)
		return err
	}

	return nil
}
