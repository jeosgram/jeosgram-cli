package services

import (
	"fmt"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/jeosgram/jeosgram-cli/utils"
	"github.com/pterm/pterm"
)

type ScreenService interface {
	ShowBusySpinner(text ...any) func(...any)
	InputCredentials() (string, string)
	InputOtp() string
}

type Terminal struct {
	DisableSpinner bool
}

func (service Terminal) isErrorExit(err error) {
	if err != nil {
		os.Exit(0)
	}
}

func (service Terminal) ShowBusySpinner(text ...any) func(...any) {
	if service.DisableSpinner {
		fmt.Println(text...)
		return func(a ...any) {}
	}

	spinner, _ := pterm.DefaultSpinner.
		WithRemoveWhenDone(true).
		WithDelay(60 * time.Millisecond).
		Start(text...)

	return spinner.Success
}

func (service Terminal) InputCredentials() (string, string) {
	var username, password string

	prompt1 := survey.Input{
		Message: "Please enter your username",
	}
	service.isErrorExit(survey.AskOne(&prompt1, &username, survey.WithValidator(utils.CheckUsername)))

	prompt2 := survey.Password{
		Message: "Please enter your password",
	}
	service.isErrorExit(survey.AskOne(&prompt2, &password, survey.WithValidator(utils.CheckPassword)))

	return utils.SanityUsername(username), password
}

func (service Terminal) InputOtp() string {
	var code string
	prompt1 := survey.Input{
		Message: "Please enter a OTP code",
	}
	service.isErrorExit(survey.AskOne(&prompt1, &code, survey.WithValidator(utils.CheckOTp)))
	return code
}
