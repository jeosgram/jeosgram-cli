package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockTerminalService struct {
	mock.Mock
}

func (m *MockTerminalService) ShowBusySpinner(text ...any) func(...any) {
	args := m.Called(text...)
	return args.Get(0).(func(...any))
}

func (m *MockTerminalService) InputCredentials() (string, string) {
	args := m.Called()
	return args.String(0), args.String(1)
}

func (m *MockTerminalService) InputOtp() string {
	args := m.Called()
	return args.String(0)
}
