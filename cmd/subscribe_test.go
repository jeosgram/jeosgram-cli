package cmd

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jeosgram/jeosgram-cli/mocks"
)

// update when the error response from function call makes something different
func Test_SubscribeCommand(t *testing.T) {

	tests := []struct {
		name       string
		prefix     string
		max        string
		deviceId   string
		until      string
		filter     string
		shouldFail bool
	}{
		{
			name:     "Run with prefix",
			prefix:   "test",
			deviceId: "we2323zdsd",
			until:    "hello",
		},
		{
			name:     "Run with filters",
			prefix:   "test",
			deviceId: "we2323zdsd",
			until:    "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			terminalService := new(mocks.MockTerminalService)

			terminalService.On("ShowBusySpinner", mock.Anything, mock.Anything).Return(func(a ...any) {})

			jeosgramApi := new(mocks.MockJeosgramAPI)

			if tt.shouldFail {
				jeosgramApi.On("EventStream", tt.deviceId, tt.prefix, mock.Anything).Return(errors.New("Error function"))
			} else {
				jeosgramApi.On("EventStream", tt.deviceId, tt.prefix, mock.Anything).Return(nil)
			}

			cmd := NewSubscribeCmd(jeosgramApi, terminalService)
			cmd.SetArgs([]string{tt.prefix})
			cmd.Flags().Set("device", tt.deviceId)
			cmd.Flags().Set("until", tt.until)
			cmd.Flags().Set("max", tt.max)

			_, err := cmd.ExecuteC()

			if tt.shouldFail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			jeosgramApi.AssertCalled(t, "EventStream", tt.deviceId, tt.prefix, mock.Anything)
			jeosgramApi.AssertNumberOfCalls(t, "EventStream", 1)
		})
	}

}
