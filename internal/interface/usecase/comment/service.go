package comment

import "go_service/internal/shared/utils"

type Service interface {
	// ListComment(sess *utils.Session, request *Request) (list ListCommentResponse, err error)
	GetCommentByOrgName(sess *utils.Session, request *CommentRequest) (listResponse CommentsResponse, err error)
	PostNewComment(sess *utils.Session, request *PostCommentRequest) (response Comment, err error)
	DeleteComments(sess *utils.Session, request *CommentRequest) (response Comment, err error)
}
