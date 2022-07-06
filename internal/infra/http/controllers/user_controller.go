package controllers

import (
	"github.com/test_server/internal/app"
	"github.com/test_server/internal/infra/http/resources"
	"github.com/test_server/internal/infra/http/validators"
	"log"
	"net/http"
)

type UserController struct {
	service           *app.UserService
	passwordValidator *validators.ResetPasswordValidator
	userValidator     *validators.UserValidator
}

func NewUserController(s *app.UserService) *UserController {
	return &UserController{
		service:           s,
		passwordValidator: validators.NewResetPasswordValidator(),
		userValidator:     validators.NewUserValidator(),
	}
}

func (u *UserController) PaginateAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userAllPageRequest, err := (*u.userValidator).ValidationUserPaginateAll(r)
		if err != nil {
			validationError(w, err)
			return
		}
		userAll, err := (*u.service).PaginateAll(uint(userAllPageRequest.Page), uint(userAllPageRequest.PageSize))

		if err != nil {
			log.Printf("UserController.GetAll(): %s", err)
			internalServerError(w, err)
			return
		}
		err = success(w, resources.MapUserDomainToDtoCollection(userAll))
		if err != nil {
			log.Printf("UserController.GetAll(): %s", err)
		}
		return
	}
}

func (u *UserController) GetOneByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID, err := (*u.userValidator).ValidationUserGetOne(r)
		if err != nil {
			validationError(w, err)
			return
		}
		userDomain, err := (*u.service).GetOneByID(userID.ID)
		if err != nil {
			internalServerError(w, err)
			return
		}
		err = success(w, resources.MapUserDomainToDto(userDomain))
		if err != nil {
			log.Printf("UserController.GetOne(): %s", err)
		}
		return
	}

}

func (u *UserController) ResetPassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := (*u.passwordValidator).ValidateAndMap(r)
		if err != nil {
			validationError(w, err)
			return
		}

		err = (*u.service).ResetPassword(user.ID, user.Password)
		if err != nil {
			internalServerError(w, err)
			return
		}

		noContent(w)
	}
}

func (u *UserController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user, err := (*u.userValidator).ValidationUserUpdate(r)
		if err != nil {
			validationError(w, err)
			return
		}

		err = (*u.service).Update(user)
		if err != nil {
			internalServerError(w, err)
			return
		}
		success(w, err)
	}
}

func (u *UserController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := (*u.userValidator).ValidationUserDelete(r)
		if err != nil {
			validationError(w, err)
			return
		}
		err = (*u.service).Delete(userID.ID)
		if err != nil {
			internalServerError(w, err)
			return
		}
		success(w, err)
	}
}
