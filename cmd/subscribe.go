package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/jeosgram/jeosgram-cli/api"
	"github.com/jeosgram/jeosgram-cli/constants"
	"github.com/jeosgram/jeosgram-cli/services"
	"github.com/jeosgram/jeosgram-cli/utils"
	"github.com/spf13/cobra"
)

const subscribeEj = `jeosgram subscribe         Subscribe to all event published
jeosgram subscribe gnss    Subscribe to events starting with "gnss" from my devices`

// subscribeCmd represents the subscribe command

func NewSubscribeCmd(jeosgram api.JeosgramClient, screenService services.ScreenService) *cobra.Command {
	subscribeCmd := &cobra.Command{
		Use:     "subscribe [event]",
		Short:   "Listen to device event stream",
		Example: subscribeEj,
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// fmt.Println("subscribe called")

			prefix, _ := utils.SliceAt(args, 0)
			deviceID, _ := cmd.Flags().GetString("device")

			until, _ := cmd.Flags().GetString("until")

			max, _ := cmd.Flags().GetInt("max")
			if max < 0 {
				max = 0
			}

			filter, _ := cmd.Flags().GetString("filter")

			isShow := makeFuncFilter(filter)

			msg := makeMsgSubscribe(prefix, deviceID)
			fmt.Println(msg)

			if filter != "" {
				fmt.Printf("This command will only show the events that match: `%s`\n", filter)
			}
			if until != "" {
				fmt.Printf("This command will exit after receiving event data matching: `%s`\n", until)
			}
			if max > 0 {
				fmt.Printf("This command will exit after receiving %d event(s)...\n", max)
			}
			fmt.Println()

			stopSpinner := screenService.ShowBusySpinner("Fetching event stream...")

			eventCount := 0
			err := jeosgram.EventStream(deviceID, prefix, func(event api.JeosgramEvent) bool {

				if isShow(event.Data) {
					if eventCount == 0 {
						stopSpinner()
					}
					_ = json.NewEncoder(os.Stdout).Encode(event)
					eventCount++
				}

				if max > 0 && eventCount >= max {
					fmt.Println(eventCount, "event(s) received. Exiting...")
					return false
				}

				if until != "" && strings.Contains(event.Data, until) {
					fmt.Println("Matching event received. Exiting...")
					return false
				}

				return true
			})
			if err != nil {
				fmt.Println(constants.PError, "Error fetching event stream:", err)
				return err
			}
			return nil
		},
	}

	subscribeCmd.Flags().StringP("device", "", "", "Listen to events from this device only")
	subscribeCmd.Flags().StringP("until", "", "", "Listen until we see an event that match this data")
	subscribeCmd.Flags().IntP("max", "", 0, "Listen until we see this many events")

	subscribeCmd.Flags().StringP("filter", "", "", "Show only the events that match this data")

	return subscribeCmd
}

func makeFuncFilter(filter string) func(s string) bool {
	if filter != "" {
		re, err := regexp.Compile(filter)
		if err == nil {
			return re.MatchString
		}

	}

	return func(s string) bool { return true }
}

func makeMsgSubscribe(event, deviceID string) string {
	switch {
	default:
		fallthrough
	case event == "" && deviceID == "":
		return "Subscribing to all events from my devices"
	case event != "" && deviceID == "":
		return fmt.Sprintf("Subscribing to `%s` from my devices", event)
	case event == "" && deviceID != "":
		return fmt.Sprintf("Subscribing to all events from %s's stream", deviceID)
	case event != "" && deviceID != "":
		return fmt.Sprintf("Subscribing to `%s` from %s's stream", event, deviceID)
	}
}
