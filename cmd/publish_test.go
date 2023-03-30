package cmd

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jeosgram/jeosgram-cli/mocks"
)

// update when the error response from function call makes something different
func Test_PublishCommand(t *testing.T) {

	tests := []struct {
		name       string
		eventName  string
		eventData  string
		shouldFail bool
	}{
		{
			name:      "Publish event successfully",
			eventName: "testEvent",
			eventData: "hello",
		},
		{
			name:       "Publish event with error",
			eventName:  "testEvent",
			eventData:  "myFunction",
			shouldFail: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			terminalService := new(mocks.MockTerminalService)

			terminalService.On("ShowBusySpinner", mock.Anything, mock.Anything).Return(func(a ...any) {})

			jeosgramApi := new(mocks.MockJeosgramAPI)

			if tt.shouldFail {
				jeosgramApi.On("Publish", tt.eventName, tt.eventData).Return(errors.New("Error function"))
			} else {
				jeosgramApi.On("Publish", tt.eventName, tt.eventData).Return(nil)
			}

			cmd := NewPublishCmd(jeosgramApi, terminalService)
			cmd.SetArgs([]string{tt.eventName, tt.eventData})
			cmd.Flags().Set("encode", "hex")

			_, err := cmd.ExecuteC()

			if tt.shouldFail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			jeosgramApi.AssertCalled(t, "Publish", tt.eventName, tt.eventData)
			jeosgramApi.AssertNumberOfCalls(t, "Publish", 1)
		})
	}

}
