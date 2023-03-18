package cmd

import (
	"fmt"

	"github.com/jeosgram/jeosgram-cli/api"
	"github.com/jeosgram/jeosgram-cli/constants"
	"github.com/jeosgram/jeosgram-cli/services"
	"github.com/jeosgram/jeosgram-cli/utils"
	"github.com/spf13/cobra"
)

const publishEj = `jeosgram publish gnss/on
jeosgram publish gnss/data '{"lat": 10.456, "lon": -85.25645}'`

// publishCmd represents the publish command
func NewPublishCmd(jeosgram api.JeosgramClient, screenService services.ScreenService) *cobra.Command {
	publishCmd := &cobra.Command{
		Use:     "publish <event> [data]",
		Short:   "Publish an event to the cloud",
		Example: publishEj,
		Args:    cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			eventName, _ := utils.SliceAt(args, 0)
			eventData, _ := utils.SliceAt(args, 1)

			encode, _ := cmd.Flags().GetString("encode")
			_ = encode

			//fmt.Println("publish called", len(args), args)

			stopSpinner := screenService.ShowBusySpinner("Publishing event:", eventName) // Publishing private event: gnss
			err := jeosgram.Publish(eventName, eventData)
			stopSpinner()

			if err != nil {
				fmt.Println(constants.PError, "Error publishing event:", err)
				return
			}

			fmt.Println("Published event:", eventName)
		},
	}

	publishCmd.Flags().StringP("encode", "e", "", "your encode")

	return publishCmd
}

func checkEncode(encode string) bool {
	switch encode {
	case "hex", "base16", "base32", "base64":
		return true
	}

	return false
}
