package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/pterm/pterm"
)

const disableSpinner = false

func ShowBusySpinner(text ...any) func(...any) {
	if disableSpinner {
		fmt.Println(text...)
		return func(a ...any) {}
	}

	spinner, _ := pterm.DefaultSpinner.
		WithRemoveWhenDone(true).
		WithDelay(60 * time.Millisecond).
		Start(text...)

	return spinner.Success
}

func SliceAt[T any](s []T, i int) (T, bool) {
	var tmp T
	if i < len(s) {
		return s[i], true
	}
	return tmp, false
}

func SanityUsername(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func IsUsername(username string) bool {
	return strings.IndexByte(username, '@') == -1
}

func sanityUsername(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func isUsername(username string) bool {
	return strings.IndexByte(username, '@') == -1
}

func InputOtp() string {
	var code string
	prompt1 := survey.Input{
		Message: "Please enter a OTP code",
	}
	isErrorExit(survey.AskOne(&prompt1, &code, survey.WithValidator(checkOTp)))
	return code
}

func isErrorExit(err error) {
	if err != nil {
		os.Exit(0)
	}
}

func InputCredentials() (string, string) {
	var username, password string

	prompt1 := survey.Input{
		Message: "Please enter your username",
	}
	isErrorExit(survey.AskOne(&prompt1, &username, survey.WithValidator(checkUsername)))

	prompt2 := survey.Password{
		Message: "Please enter your password",
	}
	isErrorExit(survey.AskOne(&prompt2, &password, survey.WithValidator(checkPassword)))

	// fmt.Printf("\n\nusername=%q password=%q\n", username, password)

	return sanityUsername(username), password
}
