package mocks

import (
	"github.com/jeosgram/jeosgram-cli/types"
	"github.com/stretchr/testify/mock"
)

type MockSessionService struct {
	mock.Mock
}

func (m *MockSessionService) Clean() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockSessionService) ReadConfig() (*types.Config, error) {
	args := m.Called()
	return args.Get(0).(*types.Config), args.Error(1)
}

func (m *MockSessionService) SaveConfig(conf *types.Config) error {
	args := m.Called(conf)
	return args.Error(0)
}

func (m *MockSessionService) SaveTokens(token *types.Token) error {
	args := m.Called(token)
	return args.Error(0)
}
