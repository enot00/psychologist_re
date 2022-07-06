package middlewares

import (
	"github.com/casbin/casbin/v2"
	"github.com/test_server/internal/app"
)

type Middlewares struct {
	e  *casbin.Enforcer
	js *app.TokenService
	us *app.UserService
}

func NewMiddlewares(e *casbin.Enforcer, js *app.TokenService, us *app.UserService) *Middlewares {
	return &Middlewares{e: e, js: js, us: us}
}
