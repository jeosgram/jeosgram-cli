package cmd

import (
	"fmt"

	"github.com/jeosgram/jeosgram-cli/api"
	"github.com/jeosgram/jeosgram-cli/constants"
	"github.com/jeosgram/jeosgram-cli/services"
	"github.com/jeosgram/jeosgram-cli/utils"
	"github.com/spf13/cobra"
)

func NewFunctionCallCmd(jeosgram api.JeosgramClient, screenService services.ScreenService) *cobra.Command {
	functionCallCmd := &cobra.Command{
		Use:   "call <device> <function> [argument]",
		Short: "Call a particular function on a device",
		Args:  cobra.RangeArgs(1, 3),
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Println("functionCall called")

			deviceID, _ := utils.SliceAt(args, 0)
			funcName, _ := utils.SliceAt(args, 1)
			funcParam, _ := utils.SliceAt(args, 2)

			stopSpinner := screenService.ShowBusySpinner("Calling function:", funcName)
			value, err := jeosgram.CallFunction(deviceID, funcName, funcParam)
			stopSpinner()

			if err != nil {
				fmt.Println(constants.PInfo, err)
				return
			}

			fmt.Println("Function return:", value)
		},
	}

	return functionCallCmd

}
