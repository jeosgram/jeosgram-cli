package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"

	"github.com/jeosgram/jeosgram-cli/api"
	"github.com/jeosgram/jeosgram-cli/constants"
	"github.com/jeosgram/jeosgram-cli/services"
	"github.com/jeosgram/jeosgram-cli/session"
)

var jeosgram *api.JeosgramAPI

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jeosgram",
	Short: "A brief description of your application",
	Long:  constants.JeosgramArt,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	conf, err := session.ReadConfig()
	if err != nil {
		panic(err) // TODO(eos175) esto no debe ser asi
	}

	terminalService := &services.Terminal{DisableSpinner: false}
	sessionService := &services.FileBasedAuthentication{}
	jeosgram = api.NewJeosgramAPI(conf.AccessToken, sessionService)

	rootCmd.AddCommand(NewLoginCmd(jeosgram, terminalService, sessionService))

	functionCmd := NewFunctionCmd()
	functionCmd.AddCommand(NewFunctionCallCmd(jeosgram, terminalService))
	functionCmd.AddCommand(NewFunctionListCmd())
	rootCmd.AddCommand(functionCmd)

	variableCmd := NewVariableCmd()
	variableCmd.AddCommand(NewGetVariableCmd(jeosgram, terminalService))
	variableCmd.AddCommand(NewVariableListCmd())
	rootCmd.AddCommand(variableCmd)

	serialCmd := NewSerialCmd()
	serialCmd.AddCommand(NewSerialListCmd())
	serialCmd.AddCommand(NewSerialMonitorCmd(terminalService))
	rootCmd.AddCommand(serialCmd)

	rootCmd.AddCommand(NewPublishCmd(jeosgram, terminalService))

	rootCmd.AddCommand(NewWebhookCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.CompletionOptions.HiddenDefaultCmd = true // oculta comando `completion`

	survey.ErrorTemplate = constants.ErrorTemplate

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jeosgram.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

}
