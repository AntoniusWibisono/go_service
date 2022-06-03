package comment

import "go_service/internal/shared/utils"

type Service interface {
	GetCommentByOrgName(sess *utils.Session, request *CommentRequest) (listResponse CommentsResponse, err error)
	PostNewComment(sess *utils.Session, request *PostCommentRequest) (response Comment, err error)
	DeleteComments(sess *utils.Session, request *CommentRequest) (response Comment, err error)
	GetMemberByOrgName(sess *utils.Session, request *MemberRequest) (listResponse ListMemberResponse, err error)
}
