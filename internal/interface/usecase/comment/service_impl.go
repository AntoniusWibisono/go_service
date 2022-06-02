package comment

import (
	"go_service/internal/domain"
	"go_service/internal/shared/utils"
)

type service struct {
	commentRepo domain.CommentsRepository
	orgsRepo    domain.OrganizationsRepository
}

func NewService(commentRepo domain.CommentsRepository, orgsRepo domain.OrganizationsRepository) *service {
	return &service{
		commentRepo: commentRepo,
		orgsRepo:    orgsRepo,
	}
}

func (s *service) DeleteComments(sess *utils.Session, request *CommentRequest) (response Comment, err error) {
	organization, err := s.orgsRepo.GetOrganizationData(request.OrganizationName)

	if err != nil {
		return
	}

	comment, err := s.commentRepo.SoftDeleteCommentData(organization.ID)

	response = Comment{
		Comment: comment.Comment,
	}

	return
}

func (s *service) PostNewComment(sess *utils.Session, request *PostCommentRequest) (response Comment, err error) {
	organization, err := s.orgsRepo.GetOrganizationData(request.OrganizationName)

	if err != nil {
		return
	}

	comment, err := s.commentRepo.PostNewComment(
		organization.ID,
		request.MemberId,
		request.Comment,
	)

	response = Comment{
		Comment: comment.Comment,
	}
	return
}

func (s *service) GetCommentByOrgName(sess *utils.Session, request *CommentRequest) (listResponse CommentsResponse, err error) {
	organization, err := s.orgsRepo.GetOrganizationData(request.OrganizationName)

	if err != nil {
		return
	}

	comments, err := s.commentRepo.GetCommentByOrganizationId(organization.ID)

	if err != nil {
		return
	}

	response := make([]Comment, 0)

	for _, comment := range comments {
		resp := Comment{
			Comment: comment.Comment,
		}
		response = append(response, resp)
	}

	listResponse = CommentsResponse{
		Comments: response,
	}

	return

}

// func (s *service) ListComment(sess *utils.Session, request *Request) (list ListCommentResponse, err error) {
// 	comments, count, err := s.commentRepo.GetListComments(
// 		request.SearchBy,
// 		request.SearchValue,
// 		request.SortBy,
// 		request.SortType,
// 		request.Page,
// 		request.PerPage,
// 	)

// 	response := make([]Comment, 0)

// 	for _, comment := range comments {
// 		resp := Comment{
// 			Comment: comment.Comment,
// 		}
// 		response = append(response, resp)
// 	}

// 	list = ListCommentResponse{
// 		Comments: response,
// 		Page:     request.Page,
// 		PerPage:  request.PerPage,
// 		Count:    count,
// 	}
// 	return
// }
