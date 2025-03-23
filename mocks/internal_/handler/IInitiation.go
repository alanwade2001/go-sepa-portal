// Code generated by mockery v2.52.3. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"

	routing "github.com/alanwade2001/go-sepa-infra/routing"
)

// IInitiation is an autogenerated mock type for the IInitiation type
type IInitiation struct {
	mock.Mock
}

// GetInitiation provides a mock function with given fields: c
func (_m *IInitiation) GetInitiation(c *gin.Context) {
	_m.Called(c)
}

// GetInitiationByID provides a mock function with given fields: c
func (_m *IInitiation) GetInitiationByID(c *gin.Context) {
	_m.Called(c)
}

// PutInitiationAccept provides a mock function with given fields: c
func (_m *IInitiation) PutInitiationAccept(c *gin.Context) {
	_m.Called(c)
}

// PutInitiationApprove provides a mock function with given fields: c
func (_m *IInitiation) PutInitiationApprove(c *gin.Context) {
	_m.Called(c)
}

// PutInitiationCancel provides a mock function with given fields: c
func (_m *IInitiation) PutInitiationCancel(c *gin.Context) {
	_m.Called(c)
}

// PutInitiationReject provides a mock function with given fields: c
func (_m *IInitiation) PutInitiationReject(c *gin.Context) {
	_m.Called(c)
}

// Register provides a mock function with given fields: r
func (_m *IInitiation) Register(r *routing.Router) {
	_m.Called(r)
}

// NewIInitiation creates a new instance of IInitiation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIInitiation(t interface {
	mock.TestingT
	Cleanup(func())
}) *IInitiation {
	mock := &IInitiation{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
