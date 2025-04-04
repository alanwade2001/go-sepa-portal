// Code generated by mockery v2.52.3. DO NOT EDIT.

package handler

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"

	routing "github.com/alanwade2001/go-sepa-infra/routing"
)

// MockIInitiation is an autogenerated mock type for the IInitiation type
type MockIInitiation struct {
	mock.Mock
}

type MockIInitiation_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIInitiation) EXPECT() *MockIInitiation_Expecter {
	return &MockIInitiation_Expecter{mock: &_m.Mock}
}

// GetInitiation provides a mock function with given fields: c
func (_m *MockIInitiation) GetInitiation(c *gin.Context) {
	_m.Called(c)
}

// MockIInitiation_GetInitiation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInitiation'
type MockIInitiation_GetInitiation_Call struct {
	*mock.Call
}

// GetInitiation is a helper method to define mock.On call
//   - c *gin.Context
func (_e *MockIInitiation_Expecter) GetInitiation(c interface{}) *MockIInitiation_GetInitiation_Call {
	return &MockIInitiation_GetInitiation_Call{Call: _e.mock.On("GetInitiation", c)}
}

func (_c *MockIInitiation_GetInitiation_Call) Run(run func(c *gin.Context)) *MockIInitiation_GetInitiation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *MockIInitiation_GetInitiation_Call) Return() *MockIInitiation_GetInitiation_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockIInitiation_GetInitiation_Call) RunAndReturn(run func(*gin.Context)) *MockIInitiation_GetInitiation_Call {
	_c.Run(run)
	return _c
}

// GetInitiationByID provides a mock function with given fields: c
func (_m *MockIInitiation) GetInitiationByID(c *gin.Context) {
	_m.Called(c)
}

// MockIInitiation_GetInitiationByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInitiationByID'
type MockIInitiation_GetInitiationByID_Call struct {
	*mock.Call
}

// GetInitiationByID is a helper method to define mock.On call
//   - c *gin.Context
func (_e *MockIInitiation_Expecter) GetInitiationByID(c interface{}) *MockIInitiation_GetInitiationByID_Call {
	return &MockIInitiation_GetInitiationByID_Call{Call: _e.mock.On("GetInitiationByID", c)}
}

func (_c *MockIInitiation_GetInitiationByID_Call) Run(run func(c *gin.Context)) *MockIInitiation_GetInitiationByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *MockIInitiation_GetInitiationByID_Call) Return() *MockIInitiation_GetInitiationByID_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockIInitiation_GetInitiationByID_Call) RunAndReturn(run func(*gin.Context)) *MockIInitiation_GetInitiationByID_Call {
	_c.Run(run)
	return _c
}

// PutInitiationAccept provides a mock function with given fields: c
func (_m *MockIInitiation) PutInitiationAccept(c *gin.Context) {
	_m.Called(c)
}

// MockIInitiation_PutInitiationAccept_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutInitiationAccept'
type MockIInitiation_PutInitiationAccept_Call struct {
	*mock.Call
}

// PutInitiationAccept is a helper method to define mock.On call
//   - c *gin.Context
func (_e *MockIInitiation_Expecter) PutInitiationAccept(c interface{}) *MockIInitiation_PutInitiationAccept_Call {
	return &MockIInitiation_PutInitiationAccept_Call{Call: _e.mock.On("PutInitiationAccept", c)}
}

func (_c *MockIInitiation_PutInitiationAccept_Call) Run(run func(c *gin.Context)) *MockIInitiation_PutInitiationAccept_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *MockIInitiation_PutInitiationAccept_Call) Return() *MockIInitiation_PutInitiationAccept_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockIInitiation_PutInitiationAccept_Call) RunAndReturn(run func(*gin.Context)) *MockIInitiation_PutInitiationAccept_Call {
	_c.Run(run)
	return _c
}

// PutInitiationApprove provides a mock function with given fields: c
func (_m *MockIInitiation) PutInitiationApprove(c *gin.Context) {
	_m.Called(c)
}

// MockIInitiation_PutInitiationApprove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutInitiationApprove'
type MockIInitiation_PutInitiationApprove_Call struct {
	*mock.Call
}

// PutInitiationApprove is a helper method to define mock.On call
//   - c *gin.Context
func (_e *MockIInitiation_Expecter) PutInitiationApprove(c interface{}) *MockIInitiation_PutInitiationApprove_Call {
	return &MockIInitiation_PutInitiationApprove_Call{Call: _e.mock.On("PutInitiationApprove", c)}
}

func (_c *MockIInitiation_PutInitiationApprove_Call) Run(run func(c *gin.Context)) *MockIInitiation_PutInitiationApprove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *MockIInitiation_PutInitiationApprove_Call) Return() *MockIInitiation_PutInitiationApprove_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockIInitiation_PutInitiationApprove_Call) RunAndReturn(run func(*gin.Context)) *MockIInitiation_PutInitiationApprove_Call {
	_c.Run(run)
	return _c
}

// PutInitiationCancel provides a mock function with given fields: c
func (_m *MockIInitiation) PutInitiationCancel(c *gin.Context) {
	_m.Called(c)
}

// MockIInitiation_PutInitiationCancel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutInitiationCancel'
type MockIInitiation_PutInitiationCancel_Call struct {
	*mock.Call
}

// PutInitiationCancel is a helper method to define mock.On call
//   - c *gin.Context
func (_e *MockIInitiation_Expecter) PutInitiationCancel(c interface{}) *MockIInitiation_PutInitiationCancel_Call {
	return &MockIInitiation_PutInitiationCancel_Call{Call: _e.mock.On("PutInitiationCancel", c)}
}

func (_c *MockIInitiation_PutInitiationCancel_Call) Run(run func(c *gin.Context)) *MockIInitiation_PutInitiationCancel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *MockIInitiation_PutInitiationCancel_Call) Return() *MockIInitiation_PutInitiationCancel_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockIInitiation_PutInitiationCancel_Call) RunAndReturn(run func(*gin.Context)) *MockIInitiation_PutInitiationCancel_Call {
	_c.Run(run)
	return _c
}

// PutInitiationReject provides a mock function with given fields: c
func (_m *MockIInitiation) PutInitiationReject(c *gin.Context) {
	_m.Called(c)
}

// MockIInitiation_PutInitiationReject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutInitiationReject'
type MockIInitiation_PutInitiationReject_Call struct {
	*mock.Call
}

// PutInitiationReject is a helper method to define mock.On call
//   - c *gin.Context
func (_e *MockIInitiation_Expecter) PutInitiationReject(c interface{}) *MockIInitiation_PutInitiationReject_Call {
	return &MockIInitiation_PutInitiationReject_Call{Call: _e.mock.On("PutInitiationReject", c)}
}

func (_c *MockIInitiation_PutInitiationReject_Call) Run(run func(c *gin.Context)) *MockIInitiation_PutInitiationReject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *MockIInitiation_PutInitiationReject_Call) Return() *MockIInitiation_PutInitiationReject_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockIInitiation_PutInitiationReject_Call) RunAndReturn(run func(*gin.Context)) *MockIInitiation_PutInitiationReject_Call {
	_c.Run(run)
	return _c
}

// Register provides a mock function with given fields: r
func (_m *MockIInitiation) Register(r *routing.Router) {
	_m.Called(r)
}

// MockIInitiation_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type MockIInitiation_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - r *routing.Router
func (_e *MockIInitiation_Expecter) Register(r interface{}) *MockIInitiation_Register_Call {
	return &MockIInitiation_Register_Call{Call: _e.mock.On("Register", r)}
}

func (_c *MockIInitiation_Register_Call) Run(run func(r *routing.Router)) *MockIInitiation_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*routing.Router))
	})
	return _c
}

func (_c *MockIInitiation_Register_Call) Return() *MockIInitiation_Register_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockIInitiation_Register_Call) RunAndReturn(run func(*routing.Router)) *MockIInitiation_Register_Call {
	_c.Run(run)
	return _c
}

// NewMockIInitiation creates a new instance of MockIInitiation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIInitiation(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIInitiation {
	mock := &MockIInitiation{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
