package cmd

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jeosgram/jeosgram-cli/mocks"
)

// update when the error response from function call makes something different
func Test_CallFunctionCommand(t *testing.T) {

	tests := []struct {
		name         string
		deviceId     string
		functionName string
		param        string
		shouldFail   bool
	}{
		{
			name:         "Call function successful",
			deviceId:     "SD323FDSD",
			functionName: "myFunction",
			param:        "myParam",
		},
		{
			name:         "Call function with error",
			deviceId:     "SD323FDSD",
			functionName: "myFunction",
			param:        "myParam",
			shouldFail:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			terminalService := new(mocks.MockTerminalService)

			terminalService.On("ShowBusySpinner", mock.Anything, mock.Anything).Return(func(a ...any) {})

			jeosgramApi := new(mocks.MockJeosgramAPI)

			if tt.shouldFail {
				jeosgramApi.On("CallFunction", tt.deviceId, tt.functionName, tt.param).Return("", errors.New("Error function"))
			} else {
				jeosgramApi.On("CallFunction", tt.deviceId, tt.functionName, tt.param).Return("Value", nil)
			}

			cmd := NewFunctionCallCmd(jeosgramApi, terminalService)
			cmd.SetArgs([]string{tt.deviceId, tt.functionName, tt.param})

			_, err := cmd.ExecuteC()

			assert.Nil(t, err)

			jeosgramApi.AssertCalled(t, "CallFunction", tt.deviceId, tt.functionName, tt.param)
			jeosgramApi.AssertNumberOfCalls(t, "CallFunction", 1)
		})
	}

}
