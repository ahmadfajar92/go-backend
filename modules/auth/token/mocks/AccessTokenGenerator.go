// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import token "github.com/purwokertodev/go-backend/modules/auth/token"

// AccessTokenGenerator is an autogenerated mock type for the AccessTokenGenerator type
type AccessTokenGenerator struct {
	mock.Mock
}

// GenerateAccessToken provides a mock function with given fields: cl
func (_m *AccessTokenGenerator) GenerateAccessToken(cl token.Claim) <-chan token.AccessTokenResponse {
	ret := _m.Called(cl)

	var r0 <-chan token.AccessTokenResponse
	if rf, ok := ret.Get(0).(func(token.Claim) <-chan token.AccessTokenResponse); ok {
		r0 = rf(cl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan token.AccessTokenResponse)
		}
	}

	return r0
}
