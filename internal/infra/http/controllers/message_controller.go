package controllers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/test_server/internal/app"
	"github.com/test_server/internal/infra/http/resources"
	"github.com/test_server/internal/infra/http/validators"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type MessageController struct {
	service   *app.MessageService
	storage   *app.StorageService
	validator *validators.MessageValidator
}

func NewMessageController(s *app.MessageService, storage *app.StorageService) *MessageController {
	return &MessageController{
		service:   s,
		storage:   storage,
		validator: validators.NewMessageValidator(),
	}
}

func (c *MessageController) AddMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		msg, err := c.validator.ValidateAndMap(r)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		if msg.FilePath != "" {
			moved, err := c.storage.MoveToPermanent(msg.FilePath)
			if err != nil {
				log.Print(err)
				internalServerError(w, err)
				return
			}

			msg.FilePath = moved.Path
		}

		addMsg, err := (*c.service).AddMessage(msg)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		err = success(w, resources.MapChatToMessageDto(addMsg))
		if err != nil {
			log.Print(err)
		}
	}
}

func (c *MessageController) UpdateMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, err := strconv.Atoi(chi.URLParam(r, "id"))

		updMsg, err := c.validator.ValidateAndMap(r)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		_, err = (*c.service).UpdateMessage(updMsg)
		if err != nil {
			log.Print(err)
			err = internalServerError(w, err)
			return
		}

		ok(w)
	}
}

func (c *MessageController) PaginateAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page, err := strconv.Atoi(chi.URLParam(r, "page"))
		if err != nil {
			log.Print(err)
			err = internalServerError(w, err)
			return
		}
		chatID, err := strconv.ParseInt(chi.URLParam(r, "chatID"), 10, 64)
		if err != nil {
			fmt.Printf("MessageController.PaginateAll(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("MessageController.PaginateAll(): %s", err)
			}
			return
		}

		msg, err := (*c.service).PaginateAll(chatID, uint(page), 5)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		err = success(w, resources.MapChatToMessageDtoCollection(msg))
		if err != nil {
			log.Printf("MessageController.PaginateAll(): %s", err)
		}
	}
}

func (c *MessageController) DeleteMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		err = (*c.service).DeleteMessage(int64(id))
		if err != nil {
			log.Print(err)
			err = internalServerError(w, err)
			return
		}

		ok(w)
	}
}

func (c *MessageController) SaveFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20)

		file, header, err := r.FormFile("file")
		if err != nil {
			log.Printf("Error: %v", err)
			badRequest(w, err)
			return
		}

		defer file.Close()

		matched, err := regexp.MatchString(`^([abc\d]+).(png|jpeg|jpg|pdf)$`, header.Filename)
		if err != nil {
			log.Printf("Error: %v", err)
			validationError(w, err)
			return
		}

		if !matched {
			validationError(w, fmt.Errorf("bad file name"))
			return
		}

		np := strings.Split(header.Filename, ".")
		newFileName := uuid.New().String() + "." + np[1]

		savedFile, err := c.storage.SaveAsTmp(newFileName, file)

		success(w, map[string]string{"file": savedFile.Path})
	}
}
