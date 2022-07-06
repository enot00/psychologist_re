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

type PsychologistValidator struct {
	validator *validator.Validate
}

func NewPsychologistValidator() *PsychologistValidator {
	return &PsychologistValidator{validator: validator.New()}
}
func (ps *PsychologistValidator) ValidationPsychologistPaginateAll(r *http.Request) (*psychologistAllPageRequest, error) {

	var data *psychologistAllPageRequest

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		log.Printf("PsychologistValidator.ValidationPsychologistPaginateAll(): %s\n", err.Error())
		return nil, err
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("page-size"))
	if err != nil {
		log.Printf("PsychologistValidator.ValidationPsychologistPaginateAll(): %s\n", err.Error())
		return nil, err
	}
	data = MapPsychologistRequestAllPage(page, pageSize)

	err = ps.validator.Struct(data)
	if err != nil {
		log.Printf("PsychologistValidator.ValidationPsychologistPaginateAll(): %s\n", err.Error())
		return nil, err
	}

	return data, nil
}
func (ps *PsychologistValidator) ValidationPsychologistGetOne(r *http.Request) (*psychologistIDRequest, error) {

	var data *psychologistIDRequest

	userID, err := strconv.ParseInt(chi.URLParam(r, "psyId"), 10, 64)
	if err != nil {
		log.Printf("Psychologist.ValidationPsychologistGetOne(): %s\n", err.Error())
		return nil, err
	}

	data = MapPsychologistIDRequest(userID)

	err = ps.validator.Struct(data)
	if err != nil {
		log.Printf("PsychologistValidator.ValidationPsychologistGetOne(): %s\n", err.Error())
		return nil, err
	}

	return data, nil
}

func (ps *PsychologistValidator) ValidationPsychologistUpdate(r *http.Request) (*domain.Psychologist, error) {
	var data PsychologistRequest

	userID, err := strconv.ParseInt(chi.URLParam(r, "psyId"), 10, 64)

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("PsychologistValidator.ValidationPsychologistUpdate(): %s\n", err.Error())
		return nil, err
	}
	data.User.ID = userID

	err = ps.validator.Struct(data)
	if err != nil {
		log.Printf("PsychologistValidator.ValidationPsychologistUpdate(): %s\n", err.Error())
		return nil, err
	}

	return MapPsychologistRequestToDomain(&data), nil
}
