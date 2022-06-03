package comment

import (
	"go_service/internal/domain"
	"go_service/internal/domain/mocks"
	"go_service/internal/shared/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_PostNewComment(t *testing.T) {
	commentsRepo := &mocks.CommentsRepository{}

	OrgsRepo := &mocks.OrganizationsRepository{}

	inputOrganization := domain.Organizations{
		ID:   "987706a6-8c82-410b-96f2-aab5bfff25e1",
		Name: "Estron",
	}

	memberId := "c00613fa-e5cf-492d-bffa-27c9b54820af"

	inputComment := "looking for SE Asia Dev"

	OrgsRepo.On("GetOrganizationData", mock.Anything).Return(inputOrganization, nil)

	var result = domain.Comments{
		Comment: inputComment,
	}

	commentsRepo.On("PostNewComment", inputOrganization.ID, memberId, inputComment).Return(result, nil)

	service := NewService(nil, nil, nil)

	service.commentRepo = commentsRepo

	service.orgsRepo = OrgsRepo

	res, err := service.PostNewComment(utils.NewSessionRequest(), &PostCommentRequest{
		OrganizationName: inputOrganization.Name,
		MemberId:         memberId,
		Comment:          inputComment,
	})

	assert.Equal(t, result.Comment, res.Comment)
	assert.NoError(t, err)

}

func Test_GetCommentByOrgName(t *testing.T) {
	commentsRepo := &mocks.CommentsRepository{}

	OrgsRepo := &mocks.OrganizationsRepository{}

	organization := domain.Organizations{
		ID: "987706a6-8c82-410b-96f2-aab5bfff25e1",
	}

	OrgsRepo.On("GetOrganizationData", mock.Anything).Return(organization, nil)

	respComment := []domain.Comments{
		{Comment: "test"},
		{Comment: "test2"},
	}

	commentsRepo.On("GetCommentByOrganizationId", mock.Anything).Return(respComment, nil)

	service := NewService(nil, nil, nil)

	service.commentRepo = commentsRepo

	service.orgsRepo = OrgsRepo

	res, err := service.GetCommentByOrgName(utils.NewSessionRequest(), &CommentRequest{
		OrganizationName: "mockOrganization",
	})

	assert.Equal(t, "test", res.Comments[0].Comment)
	assert.Equal(t, 2, len(res.Comments))
	assert.NoError(t, err)
}

func Test_DeleteComments(t *testing.T) {
	OrgsRepo := &mocks.OrganizationsRepository{}
	CommentsRepo := &mocks.CommentsRepository{}

	organization := domain.Organizations{
		ID:   "987706a6-8c82-410b-96f2-aab5bfff25e1",
		Name: "mockOrganization",
	}

	OrgsRepo.On("GetOrganizationData", mock.Anything).Return(organization, nil)
	CommentsRepo.On("SoftDeleteCommentData", mock.Anything).Return(nil)

	service := NewService(nil, nil, nil)

	service.commentRepo = CommentsRepo

	service.orgsRepo = OrgsRepo

	res, err := service.DeleteComments(utils.NewSessionRequest(), &CommentRequest{
		OrganizationName: "mockOrganization",
	})

	assert.Equal(t, "organization comments deleted", res.Message)
	assert.NoError(t, err)

}

func Test_GetMemberByOrgName(t *testing.T) {
	OrgsRepo := &mocks.OrganizationsRepository{}

	memberRepo := &mocks.MembersRepository{}

	organization := domain.Organizations{
		ID:   "987706a6-8c82-410b-96f2-aab5bfff25e1",
		Name: "mockOrganization",
	}

	OrgsRepo.On("GetOrganizationData", mock.Anything).Return(organization, nil)

	respMember := []domain.Members{
		{Username: "Jane Doe", Followers: 250, Following: 35, AvatarUrl: "www.google.pic/02"},
		{Username: "John Doe", Followers: 100, Following: 10, AvatarUrl: "www.google.pic/01"},
	}

	memberRepo.On("GetMemberData", mock.Anything).Return(respMember, nil)

	service := NewService(nil, nil, nil)

	service.memberRepo = memberRepo

	service.orgsRepo = OrgsRepo

	res, err := service.GetMemberByOrgName(utils.NewSessionRequest(), &MemberRequest{
		OrganizationName: "mockOrganization",
	})

	assert.Equal(t, "Jane Doe", res.Members[0].Username)
	assert.Equal(t, 2, len(res.Members))
	assert.NoError(t, err)
}
