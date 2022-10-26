package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// variableGetCmd represents the variableGet command
var variableGetCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("variableGet called")

		deviceID, _ := sliceAt(args, 0)
		varName, _ := sliceAt(args, 1)

		// TODO() verificar deviceID, logitudes....

		stopSpinner := showBusySpinner("Getting variable:", varName)
		value, err := jeosgram.GetVariable(deviceID, varName)
		stopSpinner()

		if err != nil {
			fmt.Println(pInfo, err)
			return
		}

		fmt.Println("Variable value:", value)

	},
}

func init() {
	variableCmd.AddCommand(variableGetCmd)
}
