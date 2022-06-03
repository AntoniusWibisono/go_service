package comment

import (
	"go_service/internal/domain"
	"go_service/internal/shared/utils"
)

type service struct {
	commentRepo domain.CommentsRepository
	orgsRepo    domain.OrganizationsRepository
	memberRepo  domain.MembersRepository
}

func NewService(commentRepo domain.CommentsRepository, orgsRepo domain.OrganizationsRepository, memberRepo domain.MembersRepository) *service {
	return &service{
		commentRepo: commentRepo,
		orgsRepo:    orgsRepo,
		memberRepo:  memberRepo,
	}
}

func (s *service) GetMemberByOrgName(sess *utils.Session, request *MemberRequest) (listResponse ListMemberResponse, err error) {
	organization, err := s.orgsRepo.GetOrganizationData(request.OrganizationName)

	if err != nil {
		return
	}

	members, err := s.memberRepo.GetMemberData(organization.ID)

	if err != nil {
		return
	}

	response := make([]Member, 0)

	for _, member := range members {
		resp := Member{
			Username:  member.Username,
			AvatarUrl: member.AvatarUrl,
			Followers: member.Followers,
			Following: member.Following,
		}
		response = append(response, resp)
	}

	listResponse = ListMemberResponse{
		Members: response,
	}

	return

}

func (s *service) DeleteComments(sess *utils.Session, request *CommentRequest) (response BasicResponse, err error) {
	organization, err := s.orgsRepo.GetOrganizationData(request.OrganizationName)

	if err != nil {
		return
	}

	err = s.commentRepo.SoftDeleteCommentData(organization.ID)

	if err != nil {
		return
	}

	response = BasicResponse{
		Message: "organization comments deleted",
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
