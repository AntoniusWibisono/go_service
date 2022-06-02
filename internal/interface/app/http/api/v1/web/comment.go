package web

import (
	"go_service/internal/interface/usecase/comment"
	"go_service/internal/shared/constants"
	"go_service/internal/shared/utils"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentService comment.Service
}

func NewSubmissionHandler(com comment.Service) *CommentHandler {
	return &CommentHandler{
		commentService: com,
	}
}

func (h *CommentHandler) DeleteCommentHandler(ctx echo.Context) (err error) {
	var sess = ctx.Get(constants.AppSessionRequest).(*utils.Session)
	var request = new(comment.CommentRequest)
	request.OrganizationName = ctx.Param("org-name")

	if err = ctx.Bind(request); err != nil {
		return sess.ResponseInvalidRequest(ctx, err.Error())
	}

	comment, err := h.commentService.DeleteComments(sess, request)
	if err != nil {
		return sess.ResponseInternalError(ctx, err.Error())
	}

	response := utils.CreateHttpResponse(constants.StatusOK, "delete success", comment)
	return sess.ResponseOK(ctx, response)
}

func (h *CommentHandler) PostNewCommentHandler(ctx echo.Context) (err error) {
	var sess = ctx.Get(constants.AppSessionRequest).(*utils.Session)
	var request = new(comment.PostCommentRequest)
	request.OrganizationName = ctx.Param("org-name")
	request.Comment = ctx.FormValue("comment")
	request.MemberId = "53432b76-9ee6-4507-86f1-53af94b8bec4"

	if err = ctx.Bind(request); err != nil {
		return sess.ResponseInvalidRequest(ctx, err.Error())
	}

	comment, err := h.commentService.PostNewComment(sess, request)
	if err != nil {
		return sess.ResponseInternalError(ctx, err.Error())
	}

	response := utils.CreateHttpResponse(constants.StatusCreated, "posting comment success", comment)
	return sess.ResponseCreated(ctx, response)
}

func (h *CommentHandler) GetCommentByOrganizationNameHandler(ctx echo.Context) (err error) {
	var sess = ctx.Get(constants.AppSessionRequest).(*utils.Session)
	var request = new(comment.CommentRequest)
	request.OrganizationName = ctx.Param("org-name")

	if err = ctx.Bind(request); err != nil {
		return sess.ResponseInvalidRequest(ctx, err.Error())
	}

	comment, err := h.commentService.GetCommentByOrgName(sess, request)
	if err != nil {
		return sess.ResponseInternalError(ctx, err.Error())
	}

	response := utils.CreateHttpResponse(constants.StatusOK, "get data success!", comment)
	return sess.ResponseOK(ctx, response)
}

// func (h *CommentHandler) ListCommentHandler(ctx echo.Context) (err error) {
// 	var sess = ctx.Get(constants.AppSessionRequest).(*utils.Session)
// 	var request = new(comment.Request)
// 	request.Page = cast.ToInt64(ctx.QueryParam("page"))
// 	request.PerPage = cast.ToInt64(ctx.QueryParam("perPage"))

// 	if err = ctx.Bind(request); err != nil {
// 		return sess.ResponseInvalidRequest(ctx, err.Error())
// 	}

// 	comments, err := h.commentService.ListComment(sess, request)
// 	if err != nil {
// 		return sess.ResponseInternalError(ctx, err.Error())
// 	}

// 	response := utils.CreateHttpResponse(constants.StatusOK, "success!", comments)
// 	return sess.ResponseOK(ctx, response)

// }
