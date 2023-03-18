package cmd

import (
	"fmt"

	"github.com/jeosgram/jeosgram-cli/api"
	"github.com/jeosgram/jeosgram-cli/constants"
	"github.com/jeosgram/jeosgram-cli/services"
	"github.com/spf13/cobra"
)

const maxRetry = 5

func NewLoginCmd(jeosgram api.JeosgramClient, screenService services.ScreenService, sessionService services.SessionService) *cobra.Command {
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Login to Jeosgram account",
		Run: func(cmd *cobra.Command, args []string) {
			const msg = "Sending login details..."

			for i := 0; i < maxRetry; i++ {
				username, password := screenService.InputCredentials()
				stopSpinner := screenService.ShowBusySpinner(msg)
				token, mfaToken, err := jeosgram.LoginByPassword(username, password)
				stopSpinner()
				if err != nil {
					fmt.Println(constants.PInfo, err)

					if err == constants.ErrRequiredMFA {
						const msg = "Sending login code..."

						for i := 0; i < maxRetry; i++ {
							otp := screenService.InputOtp()
							stopSpinner := screenService.ShowBusySpinner(msg)
							token, err = jeosgram.LoginByMFAOtp(mfaToken, otp)
							fmt.Print(err)
							stopSpinner()
							if err != nil {
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
				_ = sessionService.SaveTokens(token)
				fmt.Println(constants.PSuccess, "Successfully completed login!")
				return
			}
		},
	}

	loginCmd.Flags().StringP("username", "u", "", "your username")
	loginCmd.Flags().StringP("password", "p", "", "your password")

	loginCmd.Flags().StringP("token", "t", "", "an existing Jeosgram access token to use")

	return loginCmd
}
