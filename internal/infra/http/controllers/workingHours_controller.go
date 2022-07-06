package controllers

import (
	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/app"
	"github.com/test_server/internal/infra/http/validators"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type WorkingHoursController struct {
	service *app.WorkingHoursService
}

func NewWorkingHoursController(w *app.WorkingHoursService) *WorkingHoursController {
	return &WorkingHoursController{
		service: w,
	}
}
func (wh *WorkingHoursController) GetOneByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseInt(chi.URLParam(r, "psyID"), 10, 64)
		if err != nil {
			log.Print(err)
		}
		psy, err := (*wh.service).GetOneByID(userID)
		if err != nil {
			log.Print(err)
			err := internalServerError(w, err)
			if err != nil {
				return
			}
			return
		}
		err = success(w, psy)
		if err != nil {
			log.Printf("WorkingHoursController.GetOneById(): %s", err)
		}
		return
	}
}
func (wh *WorkingHoursController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyRead, err := ioutil.ReadAll(r.Body)

		hours, err := validators.ValidationWorkingHours(bodyRead)
		if err != nil {
			err = badRequest(w, err)
			log.Println(err)
			return
		}

		psy, err := (*wh.service).Save(&hours)
		if err != nil {
			log.Printf("WorkingHoursController.Save(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				log.Printf("WorkingHoursController.Save(): %s", err)
			}
			log.Println(err)
			return
		}

		err = success(w, psy)
		if err != nil {
			log.Printf("WorkingHoursController.Save(): %s", err)
		}
	}
}
func (wh *WorkingHoursController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		psyId, err := strconv.ParseInt(chi.URLParam(r, "psyID"), 10, 64)
		bodyRead, err := ioutil.ReadAll(r.Body)

		hours, err := validators.ValidationWorkingHours(bodyRead)
		if err != nil {
			err = badRequest(w, err)
			log.Println(err)
			return
		}

		err = (*wh.service).Update(psyId, &hours)
		if err != nil {
			log.Printf("WorkingHoursController.Update(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				log.Printf("WorkingHoursController.Update(): %s", err)
			}
			return
		}

		err = success(w, err)
		if err != nil {
			log.Printf("WorkingHoursController.Update(): %s", err)
		}
	}
}
func (wh *WorkingHoursController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		psyId, err := strconv.ParseInt(chi.URLParam(r, "psyID"), 10, 64)

		if err != nil {
			log.Print(err)
			err := badRequest(w, err)
			if err != nil {
				return
			}
			return
		}

		err = (*wh.service).Delete(psyId)
		if err != nil {
			log.Print(err)
			err = internalServerError(w, err)
			return
		}

		err = success(w, err)
		if err != nil {
			log.Print(err)
		}
	}
}
