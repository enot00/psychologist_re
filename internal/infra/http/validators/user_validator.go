package validators

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"github.com/test_server/internal/domain"
	"log"
	"net/http"
	"strconv"
)

type UserValidator struct {
	validator *validator.Validate
}

func NewUserValidator() *UserValidator {
	return &UserValidator{validator: validator.New()}
}
func (v *UserValidator) ValidationUserPaginateAll(r *http.Request) (*userAllPageRequest, error) {

	var data *userAllPageRequest

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		log.Printf("UserValidator.ValidationUserPaginateAll(): %s\n", err.Error())
		return nil, err
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("page-size"))
	if err != nil {
		log.Printf("UserValidator.ValidationUserPaginateAll(): %s\n", err.Error())
		return nil, err
	}
	data = MapUserRequestAllPage(page, pageSize)

	err = v.validator.Struct(data)
	if err != nil {
		log.Printf("UserValidator.ValidationUserPaginateAll(): %s\n", err.Error())
		return nil, err
	}

	return data, nil
}
func (v *UserValidator) ValidationUserGetOne(r *http.Request) (*userIDRequest, error) {

	var data *userIDRequest

	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		log.Printf("UserValidator.ValidationUserGetOne(): %s\n", err.Error())
		return nil, err
	}

	data = MapUserIDRequest(userID)

	err = v.validator.Struct(data)
	if err != nil {
		log.Printf("UserValidator.ValidationUserGetOne(): %s\n", err.Error())
		return nil, err
	}

	return data, nil
}

func (v *UserValidator) ValidateNewUserAndMap(r *http.Request) (*domain.User, error) {
	var data newUserRequest

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("UserValidator.ValidateNewUserAndMap(): %s\n", err.Error())
		return nil, err
	}

	err = v.validator.Struct(data)
	if err != nil {
		log.Printf("UserValidator.ValidateNewUserAndMap(): %s\n", err.Error())
		return nil, err
	}

	return MapNewUserToModel(&data), nil
}

func (v *UserValidator) ValidationUserUpdate(r *http.Request) (*domain.User, error) {
	var data userUpdateRequest

	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("UserValidator.ValidationUserUpdate(): %s\n", err.Error())
		return nil, err
	}
	data.ID = userID

	err = v.validator.Struct(data)
	if err != nil {
		log.Printf("UserValidator.ValidationUserUpdate(): %s\n", err.Error())
		return nil, err
	}

	return MapUserRequestUpdateToDomain(&data), nil
}

func (v *UserValidator) ValidationUserDelete(r *http.Request) (*userIDRequest, error) {

	var data *userIDRequest

	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		log.Printf("UserValidator.ValidationUserDelete(): %s\n", err.Error())
		return nil, err
	}

	data = MapUserIDRequest(userID)

	err = v.validator.Struct(data)
	if err != nil {
		log.Printf("UserValidator.ValidationUserDelete(): %s\n", err.Error())
		return nil, err
	}

	return data, nil
}
