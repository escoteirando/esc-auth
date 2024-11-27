package controllers

import (
	"errors"
	"fmt"

	"github.com/escoteirando/esc-auth/internal/services"
	"gofr.dev/pkg/gofr"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *gofr.Context) (interface{}, error) {
	var loginRequest LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		return nil, NewHTTPError(fmt.Sprintf("Invalid payload: %v", err), 400)
	}

	if len(loginRequest.Username) == 0 || len(loginRequest.Password) == 0 {
		return nil, errors.New("missing username or password in request body")
	}
	service := services.NewAuthService(c.SQL)
	if user, err := service.Authenticate(c.Context, loginRequest.Username, loginRequest.Password); err != nil {
		return nil, NewHTTPError("UNAUTHORIZED", 403)
	} else {
		// TODO: Retornar Token de autenticação para o
		if token, err := service.GetJWT(user); err != nil {
			return nil, NewHTTPError("UNAUTHORIZED", 403)
		} else {
			return token, nil
		}
	}
}
