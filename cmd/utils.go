package cmd

import (
	"fmt"
	"time"

	"github.com/pterm/pterm"
)

const disableSpinner = false

func showBusySpinner(text ...any) func(...any) {
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

func sliceAt[T any](s []T, i int) (T, bool) {
	var tmp T
	if i < len(s) {
		return s[i], true
	}
	return tmp, false
}
