package v1

import (
	"go_service/internal/interface/app/http/api/v1/web"
	"go_service/internal/interface/container"

	"github.com/labstack/echo/v4"
)

type Router struct {
	V1Group        *echo.Group
	CommentHandler *web.CommentHandler
}

func NewRouter(server *echo.Echo, cont *container.Container) *Router {
	return &Router{
		V1Group:        server.Group("/v1"),
		CommentHandler: web.NewSubmissionHandler(cont.CommentService),
	}
}

func (r *Router) RegisterRoutes() {
	r.V1Group.DELETE("/orgs/:org-name/comment", r.CommentHandler.DeleteCommentHandler)
	r.V1Group.POST("/orgs/:org-name/comment", r.CommentHandler.PostNewCommentHandler)
	r.V1Group.GET("/orgs/:org-name/comment", r.CommentHandler.GetCommentByOrganizationNameHandler)
	r.V1Group.GET("/orgs/:org-name/members", r.CommentHandler.GetMembersByOrganizationNameHandler)
}
