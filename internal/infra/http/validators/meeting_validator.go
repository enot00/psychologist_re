package validators

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/test_server/internal/domain"
)

type MeetingValidator struct {
	validator *validator.Validate
}

func NewMeetingValidator() *MeetingValidator {
	return &MeetingValidator{
		validator: validator.New(),
	}
}

func (m MeetingValidator) ValidateAndMap(r *http.Request, createBy string) (*domain.Meeting, error) {
	var meetingResource meetingRequest
	err := json.NewDecoder(r.Body).Decode(&meetingResource)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	if meetingResource.Endime <= meetingResource.StartTime {
		err := errors.New("end must be greater than start")
		log.Print(err)
		return nil, err
	}

	userId, err := strconv.ParseInt(r.Header.Get("user_id"), 10, 64)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	if createBy == "client" {
		meetingResource.ClientId = userId
	} else if createBy == "psychologist" {
		meetingResource.PsychologistId = userId
	}

	err = m.validator.Struct(meetingResource)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return mapMeetingRequestToDomain(&meetingResource), nil
}
