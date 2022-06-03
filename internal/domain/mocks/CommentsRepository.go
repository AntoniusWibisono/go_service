// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	domain "go_service/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// CommentsRepository is an autogenerated mock type for the CommentsRepository type
type CommentsRepository struct {
	mock.Mock
}

// GetCommentByOrganizationId provides a mock function with given fields: organizationId
func (_m *CommentsRepository) GetCommentByOrganizationId(organizationId string) ([]domain.Comments, error) {
	ret := _m.Called(organizationId)

	var r0 []domain.Comments
	if rf, ok := ret.Get(0).(func(string) []domain.Comments); ok {
		r0 = rf(organizationId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Comments)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(organizationId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostNewComment provides a mock function with given fields: organizationId, memberId, comment
func (_m *CommentsRepository) PostNewComment(organizationId string, memberId string, comment string) (domain.Comments, error) {
	ret := _m.Called(organizationId, memberId, comment)

	var r0 domain.Comments
	if rf, ok := ret.Get(0).(func(string, string, string) domain.Comments); ok {
		r0 = rf(organizationId, memberId, comment)
	} else {
		r0 = ret.Get(0).(domain.Comments)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(organizationId, memberId, comment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SoftDeleteCommentData provides a mock function with given fields: organizationId
func (_m *CommentsRepository) SoftDeleteCommentData(organizationId string) error {
	ret := _m.Called(organizationId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(organizationId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewCommentsRepositoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommentsRepository creates a new instance of CommentsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentsRepository(t NewCommentsRepositoryT) *CommentsRepository {
	mock := &CommentsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
