package http

import (
	"go_service/internal/interface/container"
	"go_service/internal/shared/constants"
	"go_service/internal/shared/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func SetupMiddlewares(server *echo.Echo, cont *container.Container) {
	server.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			reqID := ctx.Request().Header.Get("threadId")
			if len(reqID) == 0 {
				thredId := uuid.New()
				reqID = thredId.String()
			}

			conf := cont.Config
			sess := utils.NewSessionRequest()
			sess.ThreadID = reqID
			sess.AppName = conf.Name
			sess.AppVersion = conf.Version
			sess.Port = conf.HttpPort
			sess.SrcIP = ctx.RealIP()
			sess.URL = ctx.Request().URL.String()
			sess.Method = ctx.Request().Method
			sess.Header = ctx.Request().Header

			ctx.Set(constants.AppSessionRequest, sess)

			return h(ctx)
		}
	})
}
