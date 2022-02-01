package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const publishEj = `jeosgram publish gnss/on
jeosgram publish gnss/data '{"lat": 10.456, "lon": -85.25645}'`

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:     "publish <event> [data]",
	Short:   "Publish an event to the cloud",
	Example: publishEj,
	Args:    cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		eventName, _ := sliceAt(args, 0)
		eventData, _ := sliceAt(args, 1)

		encode, _ := cmd.Flags().GetString("encode")
		_ = encode

		//fmt.Println("publish called", len(args), args)

		stopSpinner := showBusySpinner("Publishing event:", eventName) // Publishing private event: gnss
		err := jeosgram.Publish(eventName, eventData)
		stopSpinner()

		if err != nil {
			fmt.Println(pError, "Error publishing event:", err)
			return
		}

		fmt.Println("Published event:", eventName)
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)

	//publishCmd.Flags().StringP("encode", "", "", "set publish encoding")
}

func checkEncode(encode string) bool {
	switch encode {
	case "hex", "base16", "base32", "base64":
		return true
	}

	return false
}
