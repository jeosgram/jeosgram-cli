package cmd

import (
	"fmt"
	"os"

	"github.com/jeosgram/jeosgram-cli/api"
	"github.com/jeosgram/jeosgram-cli/session"
	"github.com/spf13/cobra"
)

var jeosgram *api.JeosgramAPI

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jeosgram",
	Short: "A brief description of your application",
	Long:  jeosgramArt,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.CompletionOptions.HiddenDefaultCmd = true // oculta comando `completion`

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
	conf, err := session.ReadConfig()
	if err != nil {
		panic(err) // TODO(eos175) esto no debe ser asi
	}

	jeosgram = api.NewJeosgramAPI(conf.AccessToken)
}
