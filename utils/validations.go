package utils

import (
	"fmt"

	"github.com/jeosgram/jeosgram-cli/constants"
)

func checkUsername(val any) error {
	v, ok := val.(string)
	if !ok || v == "" {
		return fmt.Errorf("%s %s", constants.PError, "You need a username to log in")
	}
	return nil
}

func checkPassword(val any) error {
	v, ok := val.(string)
	if !ok || v == "" {
		return fmt.Errorf("%s %s", constants.PError, "You need a password to log in")
	}
	// verificar len
	return nil
}

func CheckOTp(val any) error {
	tmp, _ := val.(string)
	if tmp == "" {
		return fmt.Errorf("%s %s", constants.PError, "You need a OTP code to log in")
	}
	return nil
}

func CheckUsername(val any) error {
	v, ok := val.(string)
	if !ok || v == "" {
		return fmt.Errorf("%s %s", constants.PError, "You need a username to log in")
	}
	return nil
}

func CheckPassword(val any) error {
	v, ok := val.(string)
	if !ok || v == "" {
		return fmt.Errorf("%s %s", constants.PError, "You need a password to log in")
	}
	// verificar len
	return nil
}

func checkOTp(val any) error {
	tmp, _ := val.(string)
	if tmp == "" {
		return fmt.Errorf("%s %s", constants.PError, "You need a OTP code to log in")
	}
	return nil
}
