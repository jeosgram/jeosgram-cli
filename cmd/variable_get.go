package cmd

import (
	"fmt"

	"github.com/jeosgram/jeosgram-cli/api"
	"github.com/jeosgram/jeosgram-cli/constants"
	"github.com/jeosgram/jeosgram-cli/services"
	"github.com/jeosgram/jeosgram-cli/utils"
	"github.com/spf13/cobra"
)

// variableGetCmd represents the variableGet command
func NewGetVariableCmd(jeosgram api.JeosgramClient, screenService services.ScreenService) *cobra.Command {

	getVariableCmd := &cobra.Command{
		Use:   "get <device> <variableName>",
		Short: "Retrieve a value from your device(s)",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("variableGet called")

			deviceID, _ := utils.SliceAt(args, 0)
			varName, _ := utils.SliceAt(args, 1)

			// TODO() verificar deviceID, logitudes....

			stopSpinner := screenService.ShowBusySpinner("Getting variable:", varName)
			value, err := jeosgram.GetVariable(deviceID, varName)
			stopSpinner()

			if err != nil {
				fmt.Println(constants.PInfo, err)
				return err
			}

			fmt.Println("Variable value:", value)
			return nil

		},
	}

	return getVariableCmd
}
