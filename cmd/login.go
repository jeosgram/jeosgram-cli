package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/jeosgram/jeosgram-cli/api"
	"github.com/jeosgram/jeosgram-cli/session"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var errorTemplate = "{{ .Error.Error }}\n"

var (
	pInfo    = pterm.FgYellow.Sprint("!")
	pError   = pterm.FgRed.Sprint(">>")
	pSuccess = pterm.FgGreen.Sprint(">")
)

func init() {

	// cambio la plantilla de error
	survey.ErrorTemplate = errorTemplate
}

const maxRetry = 5

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Jeosgram account",
	Run: func(cmd *cobra.Command, args []string) {
		const msg = "Sending login details..."

		for i := 0; i < maxRetry; i++ {
			username, password := inputCredentials()
			stopSpinner := showBusySpinner(msg)
			token, mfaToken, err := jeosgram.LoginByPassword(username, password)
			stopSpinner()
			if err != nil {
				fmt.Println(pInfo, err)

				if err == api.ErrRequiredMFA {
					const msg = "Sending login code..."

					for i := 0; i < maxRetry; i++ {
						otp := inputOtp()
						stopSpinner := showBusySpinner(msg)
						token, err = jeosgram.LoginByMFAOtp(mfaToken, otp)
						stopSpinner()
						if err != nil {
							fmt.Println(pInfo, err)
							continue
						}
						goto end
					}

					// TODO(eos175) mostrar un error
					return
				}

				continue
			}

		end:
			_ = session.SaveTokens(token)
			fmt.Println(pSuccess, "Successfully completed login!")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// TODO() no implementado
	loginCmd.Flags().StringP("username", "u", "", "your username")
	loginCmd.Flags().StringP("password", "p", "", "your password")

	loginCmd.Flags().StringP("token", "t", "", "an existing Jeosgram access token to use")

}
