package cmd

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jeosgram/jeosgram-cli/mocks"
)

// update when the error response from function call makes something different
func Test_GetVariableCommand(t *testing.T) {

	tests := []struct {
		name         string
		deviceId     string
		variableName string
		shouldFail   bool
	}{
		{
			name:         "Get variable successfully",
			deviceId:     "as34errvdf",
			variableName: "test",
		},
		{
			name:         "Get variable with error",
			deviceId:     "as34errvdf",
			variableName: "test",
			shouldFail:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			terminalService := new(mocks.MockTerminalService)

			terminalService.On("ShowBusySpinner", mock.Anything, mock.Anything).Return(func(a ...any) {})

			jeosgramApi := new(mocks.MockJeosgramAPI)

			if tt.shouldFail {
				jeosgramApi.On("GetVariable", tt.deviceId, tt.variableName).Return("", errors.New("Error function"))
			} else {
				jeosgramApi.On("GetVariable", tt.deviceId, tt.variableName).Return("Value", nil)
			}

			cmd := NewGetVariableCmd(jeosgramApi, terminalService)
			cmd.SetArgs([]string{tt.deviceId, tt.variableName})

			_, err := cmd.ExecuteC()

			if tt.shouldFail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			jeosgramApi.AssertCalled(t, "GetVariable", tt.deviceId, tt.variableName)
			jeosgramApi.AssertNumberOfCalls(t, "GetVariable", 1)
		})
	}

}
