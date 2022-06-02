package http

import (
	"fmt"
	"go_service/internal/interface/container"
	"net/http"
	"time"

	v1 "go_service/internal/interface/app/http/api/v1"

	"github.com/labstack/echo/v4"
)

func StartHTTPServer(cont *container.Container) (err error) {
	server := echo.New()

	SetupMiddlewares(server, cont)
	v1Router := v1.NewRouter(server, cont)
	v1Router.RegisterRoutes()

	server.GET("/", func(ctx echo.Context) error {
		msg := fmt.Sprintf("service up and running... (%s)", time.Now().Format(time.RFC3339))
		return ctx.String(http.StatusOK, msg)
	})

	return server.Start(fmt.Sprintf("%s:%d", cont.Config.Address, cont.Config.HttpPort))
}
