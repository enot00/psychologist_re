package validators

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/test_server/internal/domain"
	"log"
	"net/http"
)

type MessageValidator struct {
	validator *validator.Validate
}

func NewMessageValidator() *MessageValidator {
	return &MessageValidator{
		validator: validator.New(),
	}
}

func (m MessageValidator) ValidateAndMap(r *http.Request) (*domain.Message, error) {
	var messageResource messageRequest
	err := json.NewDecoder(r.Body).Decode(&messageResource)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	err = m.validator.Struct(messageResource)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return mapMessageRequestToChat(&messageResource), nil
}
