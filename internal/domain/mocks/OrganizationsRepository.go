// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	domain "go_service/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// OrganizationsRepository is an autogenerated mock type for the OrganizationsRepository type
type OrganizationsRepository struct {
	mock.Mock
}

// GetOrganizationData provides a mock function with given fields: organizationName
func (_m *OrganizationsRepository) GetOrganizationData(organizationName string) (domain.Organizations, error) {
	ret := _m.Called(organizationName)

	var r0 domain.Organizations
	if rf, ok := ret.Get(0).(func(string) domain.Organizations); ok {
		r0 = rf(organizationName)
	} else {
		r0 = ret.Get(0).(domain.Organizations)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(organizationName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewOrganizationsRepositoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrganizationsRepository creates a new instance of OrganizationsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrganizationsRepository(t NewOrganizationsRepositoryT) *OrganizationsRepository {
	mock := &OrganizationsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
