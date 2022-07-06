package validators

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/test_server/internal/domain"
	"log"
	"net/http"
	"strconv"
)

type ResetPasswordValidator struct {
	validator *validator.Validate
}

func NewResetPasswordValidator() *ResetPasswordValidator {
	return &ResetPasswordValidator{validator: validator.New()}
}

func (v *ResetPasswordValidator) ValidateAndMap(r *http.Request) (*domain.User, error) {
	var data ResetPasswordRequest

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("UserValidator.ValidateAndMap(): %s\n", err.Error())
		return nil, err
	}

	err = v.validator.Struct(data)
	if err != nil {
		log.Printf("UserValidator.ValidateAndMap(): %s\n", err.Error())
		return nil, err
	}

	user := MapResetPasswordRequestToModel(&data)
	user.ID = id

	return user, nil
}
