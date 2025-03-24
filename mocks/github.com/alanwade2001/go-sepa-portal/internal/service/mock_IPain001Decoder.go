// Code generated by mockery v2.52.3. DO NOT EDIT.

package service

import (
	pain_001_001_03 "github.com/alanwade2001/go-sepa-iso/pain_001_001_03"
	mock "github.com/stretchr/testify/mock"
)

// MockIPain001Decoder is an autogenerated mock type for the IPain001Decoder type
type MockIPain001Decoder struct {
	mock.Mock
}

type MockIPain001Decoder_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIPain001Decoder) EXPECT() *MockIPain001Decoder_Expecter {
	return &MockIPain001Decoder_Expecter{mock: &_m.Mock}
}

// Decode provides a mock function with given fields: content
func (_m *MockIPain001Decoder) Decode(content string) (*pain_001_001_03.Document, error) {
	ret := _m.Called(content)

	if len(ret) == 0 {
		panic("no return value specified for Decode")
	}

	var r0 *pain_001_001_03.Document
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*pain_001_001_03.Document, error)); ok {
		return rf(content)
	}
	if rf, ok := ret.Get(0).(func(string) *pain_001_001_03.Document); ok {
		r0 = rf(content)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pain_001_001_03.Document)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(content)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIPain001Decoder_Decode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Decode'
type MockIPain001Decoder_Decode_Call struct {
	*mock.Call
}

// Decode is a helper method to define mock.On call
//   - content string
func (_e *MockIPain001Decoder_Expecter) Decode(content interface{}) *MockIPain001Decoder_Decode_Call {
	return &MockIPain001Decoder_Decode_Call{Call: _e.mock.On("Decode", content)}
}

func (_c *MockIPain001Decoder_Decode_Call) Run(run func(content string)) *MockIPain001Decoder_Decode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockIPain001Decoder_Decode_Call) Return(_a0 *pain_001_001_03.Document, _a1 error) *MockIPain001Decoder_Decode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIPain001Decoder_Decode_Call) RunAndReturn(run func(string) (*pain_001_001_03.Document, error)) *MockIPain001Decoder_Decode_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIPain001Decoder creates a new instance of MockIPain001Decoder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIPain001Decoder(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIPain001Decoder {
	mock := &MockIPain001Decoder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
