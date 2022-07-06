package controllers

import (
	"github.com/test_server/internal/app"
	"github.com/test_server/internal/infra/http/resources"
	"github.com/test_server/internal/infra/http/validators"
	"log"
	"net/http"
)

type PsychologistsController struct {
	service                *app.PsychologistService
	psychologistsValidator *validators.PsychologistValidator
}

func NewPsychologistsController(s *app.PsychologistService) *PsychologistsController {
	return &PsychologistsController{
		service:                s,
		psychologistsValidator: validators.NewPsychologistValidator(),
	}
}

func (ps *PsychologistsController) PaginateAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		psyAllPageRequest, _ := (*ps.psychologistsValidator).ValidationPsychologistPaginateAll(r)
		psyAll, err := (*ps.service).PaginateAll(uint(psyAllPageRequest.Page), uint(psyAllPageRequest.PageSize))
		if err != nil {
			log.Printf("PsychologistsController.PaginateAll(): %s", err)
			internalServerError(w, err)
			return
		}
		err = success(w, resources.MapPsychologistDomainToDtoCollection(psyAll))
		if err != nil {
			log.Printf("PsychologistsController.GetAll(): %s", err)
		}
		return
	}
}

func (ps *PsychologistsController) GetOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		psyId, _ := (*ps.psychologistsValidator).ValidationPsychologistGetOne(r)
		psyOne, err := (*ps.service).GetOne(uint(psyId.ID))
		if err != nil {
			log.Printf("PsychologistsController.One(): %s", err)
			internalServerError(w, err)
			return
		}
		err = success(w, resources.MapPsychologistDomainToDto(psyOne))
		if err != nil {
			log.Printf("PsychologistsController.GetOne(): %s", err)
			return
		}
		return
	}

}

func (ps *PsychologistsController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		psy, _ := (*ps.psychologistsValidator).ValidationPsychologistUpdate(r)

		err := (*ps.service).Update(psy)
		if err != nil {
			internalServerError(w, err)
			return
		}
		success(w, err)
	}
}
