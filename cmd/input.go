package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func isErrorExit(err error) {
	if err != nil {
		os.Exit(0)
	}
}

func sanityUsername(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func isUsername(username string) bool {
	return strings.IndexByte(username, '@') == -1
}

// ----------------------------------

/*

	https://github.com/AlecAivazis/survey#password

*/

func inputCredentials() (string, string) {
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

func inputOtp() string {
	var code string
	prompt1 := survey.Input{
		Message: "Please enter a OTP code",
	}
	isErrorExit(survey.AskOne(&prompt1, &code, survey.WithValidator(checkOTp)))
	return code
}

// ---------------

func checkUsername(val any) error {
	v, ok := val.(string)
	if !ok || v == "" {
		return fmt.Errorf("%s %s", pError, "You need a username to log in")
	}
	return nil
}

func checkPassword(val any) error {
	v, ok := val.(string)
	if !ok || v == "" {
		return fmt.Errorf("%s %s", pError, "You need a password to log in")
	}
	// verificar len
	return nil
}

func checkOTp(val any) error {
	tmp, _ := val.(string)
	if tmp == "" {
		return fmt.Errorf("%s %s", pError, "You need a OTP code to log in")
	}
	return nil
}
