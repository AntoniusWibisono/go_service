package utils

import (
	"go_service/internal/shared/constants"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Session struct {
	ThreadID                string
	AppName, AppVersion, IP string
	Port                    int
	SrcIP, URL, Method      string
	Header, Request         interface{}
}

func NewSessionRequest() *Session {
	return &Session{
		Header:  map[string]interface{}{},
		Request: struct{}{},
	}
}

func (session *Session) ResponseOK(ctx echo.Context, response interface{}) error {
	return ctx.JSON(http.StatusOK, response)
}

func (session *Session) ResponseCreated(ctx echo.Context, response interface{}) error {
	return ctx.JSON(http.StatusCreated, response)
}

func (session *Session) ResponseInternalError(ctx echo.Context, message string) error {
	response := CreateHttpResponse(constants.StatusInternalError, message, nil)
	return ctx.JSON(http.StatusOK, response)
}

func (session *Session) ResponseInvalidRequest(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusOK, CreateHttpResponse(constants.StatusInvalidRequest, message, nil))
}
