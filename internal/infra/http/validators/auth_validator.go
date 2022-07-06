package validators

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/test_server/internal/domain"
	"log"
	"net/http"
)

type AuthValidator struct {
	validator *validator.Validate
}

func NewAuthValidator() *AuthValidator {
	return &AuthValidator{validator: validator.New()}
}

func (v *AuthValidator) ValidateAndMap(r *http.Request) (*domain.User, error) {
	var data authRequest

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("UserValidator.ValidateAndMap(): %s\n", err.Error())
		return nil, err
	}

	err = v.validator.Struct(data)
	if err != nil {
		log.Printf("UserValidator.ValidateAndMap(): %s\n", err.Error())
		return nil, err
	}

	return MapAuthDataToModel(&data), nil
}
