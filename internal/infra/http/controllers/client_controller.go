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

type ClientController struct {
	clientService *app.ClientService
	validator     *validators.ClientValidator
}

func NewClientController(s *app.ClientService) *ClientController {
	return &ClientController{
		clientService: s,
		validator:     validators.NewClientValidator(),
	}
}

func (c *ClientController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clnt, err := c.validator.ValidateAndMap(r)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		svdClnt, err := (*c.clientService).SaveClient(clnt)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		err = success(w, resources.MapDomainToClientDto(svdClnt))
		if err != nil {
			log.Print(err)
		}
	}
}

func (c *ClientController) PaginateAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page, err := strconv.Atoi(chi.URLParam(r, "page"))
		if err != nil {
			log.Print(err)
			err = internalServerError(w, err)
			return
		}
		clnts, err := (*c.clientService).PaginateAllClients(uint(page), 5)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		err = success(w, resources.MapDomainToClientDtoCollection(clnts))
		if err != nil {
			log.Print(err)
		}
	}
}

func (c *ClientController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}
		clnt, err := (*c.clientService).FindOneClient(id)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		err = success(w, resources.MapDomainToClientDto(clnt))
		if err != nil {
			log.Print(err)
		}
	}
}

func (c *ClientController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clnt, err := c.validator.ValidateAndMap(r)
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
		clnt.Id = id

		err = (*c.clientService).UpdateClient(clnt)
		if err != nil {
			log.Print(err)
			err = internalServerError(w, err)
			return
		}

		ok(w)
	}
}

func (c *ClientController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		err = (*c.clientService).DeleteClient(int64(id))
		if err != nil {
			log.Print(err)
			err = internalServerError(w, err)
			return
		}

		ok(w)
	}
}
