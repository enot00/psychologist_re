package controllers

import (
	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/app"
	"github.com/test_server/internal/infra/http/validators"
	"log"
	"net/http"
	"strconv"
)

type ChatController struct {
	service   *app.ChatService
	validator *validators.ChatValidator
}

func NewChatController(s *app.ChatService) *ChatController {
	return &ChatController{service: s, validator: validators.NewChatValidator()}
}

func (c *ChatController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dto, err := validators.MapChatJsonToDto(r.Body)
		if err != nil {
			badRequest(w, err)
			return
		}

		err = c.validator.Validate(dto)
		if err != nil {
			badRequest(w, err)
			return
		}

		dto, err = c.service.Add(dto)
		if err != nil {
			internalServerError(w, err)
			return
		}

		err = success(w, dto)
		if err != nil {
			log.Println(err)
		}
	}
}

func (c *ChatController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "chatId"), 10, 64)
		if err != nil {
			log.Println(err)
			badRequest(w, err)
			return
		}

		dto, err := (*c.service).FindChat(id)
		if err != nil {
			log.Println(err)
			internalServerError(w, err)
			return
		}

		err = success(w, dto)
		if err != nil {
			log.Println(err)
		}
	}
}

func (c *ChatController) FindUserChats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "userId"), 10, 64)
		if err != nil {
			log.Println(err)
			badRequest(w, err)
			return
		}

		dto, err := (*c.service).FindUserChats(id)
		if err != nil {
			log.Println(err)
			internalServerError(w, err)
			return
		}

		err = success(w, dto)
		if err != nil {
			log.Println(err)
		}
	}
}

func (c *ChatController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "chatId"), 10, 64)
		if err != nil {
			log.Println(err)
			badRequest(w, err)
			return
		}

		err = (*c.service).Delete(id)
		if err != nil {
			log.Println(err)
			internalServerError(w, err)
			return
		}

		err = success(w, "ok")
		if err != nil {
			log.Println(err)
		}
	}
}
