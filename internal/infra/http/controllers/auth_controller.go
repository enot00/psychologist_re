package controllers

import (
	"errors"
	"fmt"
	"github.com/test_server/internal/app"
	"github.com/test_server/internal/domain"
	appErrors "github.com/test_server/internal/errors"
	"github.com/test_server/internal/infra/http/resources"
	"github.com/test_server/internal/infra/http/validators"
	"net/http"
	"time"
)

type ApiAuthenticationController struct {
	authValidator *validators.AuthValidator
	userValidator *validators.UserValidator
	service       *app.AuthenticationService
}

func NewApiAuthenticationController(a *app.AuthenticationService) *ApiAuthenticationController {
	return &ApiAuthenticationController{
		authValidator: validators.NewAuthValidator(),
		userValidator: validators.NewRegistrationValidator(),
		service:       a,
	}
}

func (c *ApiAuthenticationController) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := (*c.userValidator).ValidateAndMap(r)
		if err != nil {
			validationError(w, err)
			return
		}

		saved, tokens, err := (*c.service).SignUp(user)
		if err != nil {
			internalServerError(w, err)
			return
		}

		c.setTokens(w, r.Host, tokens)

		created(w, resources.MapModelToUserResource(saved))
	}
}

func (c *ApiAuthenticationController) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := (*c.authValidator).ValidateAndMap(r)
		if err != nil {
			validationError(w, err)
			return
		}

		logged, tokens, err := (*c.service).Login(user.Email, user.Password)
		if err != nil {
			if errors.Is(err, appErrors.ErrAuthenticationFailed) {
				unauthorizedError(w, err)
				return
			}

			internalServerError(w, err)
			return
		}

		c.setTokens(w, r.Host, tokens)

		success(w, resources.MapModelToUserResource(logged))
	}
}

func (c *ApiAuthenticationController) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var refreshToken domain.Token

		accessToken := r.Context().Value("access_token")

		refreshCookie, err := r.Cookie("refresh_token")
		if !errors.Is(err, http.ErrNoCookie) {
			refreshToken.TokenString = refreshCookie.Value
		}

		err = (*c.service).Logout(accessToken.(domain.Token), refreshToken)
		if err != nil {
			internalServerError(w, err)
			return
		}

		c.resetRefreshCookie(w)

		noContent(w)
	}
}

func (c *ApiAuthenticationController) Refresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("refresh_token")
		if errors.Is(err, http.ErrNoCookie) {
			badRequest(w, err)
			return
		}

		tokens, err := (*c.service).RefreshTokens(cookie.Value)
		if err != nil {
			internalServerError(w, err)
			return
		}

		c.setTokens(w, r.Host, tokens)

		noContent(w)
	}
}

func (c *ApiAuthenticationController) setTokens(w http.ResponseWriter, domain string, tokens *domain.Tokens) {
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", tokens.Access.TokenString))

	refreshCookie := &http.Cookie{
		Name:  "refresh_token",
		Value: tokens.Refresh.TokenString,
		//Domain:   domain, // not working on localhost
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   3600 * 24 * 14,
		Path:     "/",
	}

	http.SetCookie(w, refreshCookie)
}

func (c *ApiAuthenticationController) resetRefreshCookie(w http.ResponseWriter) {
	refreshCookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now(),
		Path:     "/",
	}

	http.SetCookie(w, refreshCookie)
}
