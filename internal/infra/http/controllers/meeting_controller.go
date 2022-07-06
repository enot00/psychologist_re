package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/app"
	"github.com/test_server/internal/infra/http/resources"
	"github.com/test_server/internal/infra/http/validators"
)

type MeetingController struct {
	meetingService *app.MeetingService
	validator      *validators.MeetingValidator
}

func NewMeetingController(s *app.MeetingService) *MeetingController {
	return &MeetingController{
		meetingService: s,
		validator:      validators.NewMeetingValidator(),
	}
}

func (c *MeetingController) CreateByPsychologist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mtng, err := c.validator.ValidateAndMap(r, "psychologist")
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		svdMtng, err := (*c.meetingService).SaveMeetingByPsychologist(mtng)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		err = success(w, resources.MapDomainToMeetingDto(svdMtng))
		if err != nil {
			log.Print(err)
		}
	}
}

func (c *MeetingController) CreateByClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mtng, err := c.validator.ValidateAndMap(r, "client")
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		svdMtng, err := (*c.meetingService).SaveMeetingByClient(mtng)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		err = success(w, resources.MapDomainToMeetingDto(svdMtng))
		if err != nil {
			log.Print(err)
		}
	}
}

func (c *MeetingController) PaginateAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page, err := strconv.Atoi(chi.URLParam(r, "page"))
		if err != nil {
			log.Print(err)
			err = internalServerError(w, err)
			return
		}
		mtngs, err := (*c.meetingService).PaginateAllMeetings(uint(page), 5)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		err = success(w, resources.MapDomainToMeetingDtoCollection(mtngs))
		if err != nil {
			log.Print(err)
		}
	}
}

func (c *MeetingController) PaginateAllByPsychologist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		psychologistId, err := strconv.ParseInt(chi.URLParam(r, "psychologist_id"), 10, 64)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		page, err := strconv.Atoi(chi.URLParam(r, "page"))
		if err != nil {
			log.Print(err)
			err = internalServerError(w, err)
			return
		}
		mtngs, err := (*c.meetingService).PaginateAllPsychologistMeetings(
			psychologistId, uint(page), 5,
		)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		err = success(w, resources.MapDomainToMeetingDtoCollection(mtngs))
		if err != nil {
			log.Print(err)
		}
	}
}

func (c *MeetingController) PaginateAllByClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientId, err := strconv.ParseInt(chi.URLParam(r, "client_id"), 10, 64)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		page, err := strconv.Atoi(chi.URLParam(r, "page"))
		if err != nil {
			log.Print(err)
			err = internalServerError(w, err)
			return
		}
		mtngs, err := (*c.meetingService).PaginateAllClientMeetings(
			clientId, uint(page), 5,
		)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		err = success(w, resources.MapDomainToMeetingDtoCollection(mtngs))
		if err != nil {
			log.Print(err)
		}
	}
}

func (c *MeetingController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}
		clnt, err := (*c.meetingService).FindOneMeeting(id)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		err = success(w, resources.MapDomainToMeetingDto(clnt))
		if err != nil {
			log.Print(err)
		}
	}
}

func (c *MeetingController) UpdateByPsychologist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mtng, err := c.validator.ValidateAndMap(r, "psychologist")
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}
		mtng.Id = id

		err = (*c.meetingService).UpdateMeetingByPsychologist(mtng)
		if err != nil {
			log.Print(err)
			err = internalServerError(w, err)
			return
		}

		ok(w)
	}
}

func (c *MeetingController) UpdateByClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mtng, err := c.validator.ValidateAndMap(r, "client")
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}
		mtng.Id = id

		err = (*c.meetingService).UpdateMeetingByClient(mtng)
		if err != nil {
			log.Print(err)
			err = internalServerError(w, err)
			return
		}

		ok(w)
	}
}

func (c *MeetingController) DeleteByPsychologist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		meetingId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		psychologistId, err := strconv.ParseInt(r.Header.Get("user_id"), 10, 64)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		err = (*c.meetingService).DeleteMeetingByPsychologist(meetingId, psychologistId)
		if err != nil {
			log.Print(err)
			err = internalServerError(w, err)
			return
		}

		ok(w)
	}
}

func (c *MeetingController) DeleteByClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		meetingId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		clientId, err := strconv.ParseInt(r.Header.Get("user_id"), 10, 64)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		err = (*c.meetingService).DeleteMeetingByClient(meetingId, clientId)
		if err != nil {
			log.Print(err)
			err = internalServerError(w, err)
			return
		}

		ok(w)
	}
}
