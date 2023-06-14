package cmd

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/jeosgram/jeosgram-cli/constants"
	"github.com/jeosgram/jeosgram-cli/mocks"
	"github.com/jeosgram/jeosgram-cli/types"
	"github.com/stretchr/testify/assert"
)

func Test_LoginCommand(t *testing.T) {

	tests := []struct {
		name        string
		username    string
		password    string
		otp         string
		mfaToken    string
		token       types.Token
		failLogin   bool
		requiresMFA bool
		failMFA     bool
	}{
		{
			name:     "Standard login success",
			username: "Test user",
			password: "User password",
			token: types.Token{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
			},
			requiresMFA: false,
		},
		{
			name:     "Standard login retry if it fails",
			username: "Test user",
			password: "User password",
			token: types.Token{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
			},
			failLogin:   true,
			requiresMFA: false,
		},
		{
			name:     "Login with MFA",
			username: "Test user",
			password: "User password",
			otp:      "012023",
			mfaToken: "so4545dfdfwe",
			token: types.Token{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
			},
			requiresMFA: true,
			failMFA:     false,
		},
		{
			name:     "Login with MFA retry if it fails",
			username: "Test user",
			password: "User password",
			otp:      "012023",
			mfaToken: "so4545dfdfwe",
			token: types.Token{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
			},
			requiresMFA: true,
			failMFA:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			terminalService := new(mocks.MockTerminalService)
			terminalService.On("InputCredentials").Return(tt.username, tt.password)
			terminalService.On("ShowBusySpinner", mock.Anything).Return(func(a ...any) {})
			terminalService.On("InputOtp").Return(tt.otp)

			sessionService := new(mocks.MockSessionService)
			sessionService.On("SaveTokens", &tt.token).Return(nil)

			jeosgramApi := new(mocks.MockJeosgramAPI)

			if tt.failLogin {
				jeosgramApi.On("LoginByPassword", tt.username, tt.password).Return(&types.Token{}, "", errors.New("Test error"))
			}

			if tt.requiresMFA {
				jeosgramApi.On("LoginByPassword", tt.username, tt.password).Return(&tt.token, tt.mfaToken, constants.ErrRequiredMFA)
			} else {
				jeosgramApi.On("LoginByPassword", tt.username, tt.password).Return(&tt.token, tt.mfaToken, nil)
			}

			if tt.failMFA {
				jeosgramApi.On("LoginByMFAOtp", tt.mfaToken, tt.otp).Return(&types.Token{}, errors.New("Test error"))
			} else {
				jeosgramApi.On("LoginByMFAOtp", tt.mfaToken, tt.otp).Return(&tt.token, nil)
			}

			cmd := NewLoginCmd(jeosgramApi, terminalService, sessionService)

			_, err := cmd.ExecuteC()

			if tt.requiresMFA && tt.failMFA {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			if tt.failLogin {
				jeosgramApi.AssertNumberOfCalls(t, "LoginByPassword", 5)
			} else {
				jeosgramApi.AssertNumberOfCalls(t, "LoginByPassword", 1)
			}

			jeosgramApi.AssertCalled(t, "LoginByPassword", tt.username, tt.password)

			if tt.requiresMFA {
				callForMFA := 1
				if tt.failMFA {
					callForMFA = 5
				}
				jeosgramApi.AssertCalled(t, "LoginByMFAOtp", tt.mfaToken, tt.otp)
				jeosgramApi.AssertNumberOfCalls(t, "LoginByMFAOtp", callForMFA)
			}
			if tt.failMFA || tt.failLogin {
				sessionService.AssertNotCalled(t, "SaveTokens", &tt.token)
			} else {
				sessionService.AssertCalled(t, "SaveTokens", &tt.token)
			}
		})
	}
}
