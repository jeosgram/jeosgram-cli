package mocks

import (
	"github.com/jeosgram/jeosgram-cli/api"
	"github.com/jeosgram/jeosgram-cli/types"
	"github.com/stretchr/testify/mock"
)

type MockJeosgramAPI struct {
	mock.Mock
}

func (m *MockJeosgramAPI) CallFunction(deviceID, funcName, funcParam string) (any, error) {
	args := m.Called(deviceID, funcName, funcParam)
	return args.Get(0), args.Error(1)
}
func (m *MockJeosgramAPI) GetVariable(deviceID, varName string) (any, error) {
	args := m.Called(deviceID, varName)
	return args.Get(0), args.Error(1)
}
func (m *MockJeosgramAPI) SignalDevice(deviceID string, signal bool) {
	m.Called(deviceID, signal)
}
func (m *MockJeosgramAPI) LoginByPassword(username, password string) (*types.Token, string, error) {
	args := m.Called(username, password)
	return args.Get(0).(*types.Token), args.String(1), args.Error(2)
}
func (m *MockJeosgramAPI) LoginByMFAOtp(mfaToken, otp string) (*types.Token, error) {
	args := m.Called(mfaToken, otp)
	return args.Get(0).(*types.Token), args.Error(1)
}
func (m *MockJeosgramAPI) LoginByRefreshToken(refreshToken string) (*types.Token, error) {
	args := m.Called(refreshToken)
	return args.Get(0).(*types.Token), args.Error(1)
}
func (m *MockJeosgramAPI) Publish(eventName, eventData string) error {
	args := m.Called(eventName, eventData)
	return args.Error(0)
}
func (m *MockJeosgramAPI) EventStream(deviceID, eventName string, fun func(event api.JeosgramEvent) bool) error {
	args := m.Called(deviceID, eventName, fun)
	return args.Error(0)
}
