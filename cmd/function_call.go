package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// functionCallCmd represents the functionCall command
var functionCallCmd = &cobra.Command{
	Use:   "call <device> <function> [argument]",
	Short: "Call a particular function on a device",
	Args:  cobra.RangeArgs(1, 3),
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("functionCall called")

		deviceID, _ := sliceAt(args, 0)
		funcName, _ := sliceAt(args, 1)
		funcParam, _ := sliceAt(args, 2)

		// TODO() verificar deviceID, logitudes....

		stopSpinner := showBusySpinner("Calling function:", funcName)
		value, err := jeosgram.CallFunction(deviceID, funcName, funcParam)
		stopSpinner()

		if err != nil {
			fmt.Println(pInfo, err)
			return
		}

		fmt.Println("Function return:", value)
	},
}

func init() {
	functionCmd.AddCommand(functionCallCmd)
}
