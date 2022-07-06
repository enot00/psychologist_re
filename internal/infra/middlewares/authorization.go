package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/test_server/internal/domain"
	appErrors "github.com/test_server/internal/errors"
	"github.com/upper/db/v4"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (m *Middlewares) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			unauthorizedError(w, errors.New("empty auth header"))
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			unauthorizedError(w, errors.New("invalid auth header"))
			return
		}

		if len(headerParts[1]) == 0 {
			unauthorizedError(w, errors.New("token is empty"))
			return
		}

		payload, err := (*m.js).Validate(headerParts[1])
		if err != nil {
			unauthorizedError(w, err)
			return
		}

		res, err := (*m.e).Enforce(strconv.Itoa(payload.Role), r.URL.Path, r.Method)

		if !res {
			unauthorizedError(w, appErrors.ErrAuthorizationFailed)
			return
		}

		user, err := (*m.us).GetOneByID(payload.UserId)
		if err != nil {
			if errors.Is(err, db.ErrNoMoreRows) {
				unauthorizedError(w, err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)

			err = json.NewEncoder(w).Encode(map[string]interface{}{"error": err})
			if err != nil {
				log.Println(err)
			}
			return
		}

		if !user.DeletedAt.IsZero() {
			unauthorizedError(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		ctx = context.WithValue(ctx, "access_token", domain.Token{
			IssuedAt:    payload.IssuedAt.Time,
			ExpiresAt:   payload.ExpiresAt.Time,
			TokenString: headerParts[1],
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func unauthorizedError(w http.ResponseWriter, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	return json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
}
