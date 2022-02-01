package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

/*

/v1/devices/events/<?prefix>
/v1/devices/<?device_id>/events/<?prefix>

*/

func makeURL(deviceID, eventName string) string {
	url := "/v1/devices"
	if deviceID != "" {
		url += "/" + deviceID
	}
	url += "/events"
	if eventName != "" {
		url += "/" + eventName
	}
	return url
}

func valueSSE(s string) (string, bool) {
	const sep = ':'
	if i := strings.IndexByte(s, sep); i != -1 {
		return strings.TrimSpace(s[i+1:]), true
	}
	return "", false
}

// -------------------------------------

type JeosgramEvent struct {
	Event  string `json:"event"`
	Data   string `json:"data,omitempty"`
	Date   string `json:"date"`
	Coreid string `json:"coreid"`
}

func (api JeosgramAPI) EventStream(deviceID, eventName string, fun func(event JeosgramEvent) bool) error {
	uri := makeURL(deviceID, eventName)

	res, err := api.get(uri)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// notificar ok

	return processStream(res.Body, fun)
}

func processStream(body io.Reader, fun func(event JeosgramEvent) bool) error {
	var eventName, eventData string

	buf := bufio.NewReader(body)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			return err
		}

		if debug {
			fmt.Printf("%q\n", line)
		}

		switch {
		case strings.HasPrefix(line, "event"):
			eventName, _ = valueSSE(line)

		case strings.HasPrefix(line, "data"):
			event := JeosgramEvent{
				Event: eventName,
			}

			eventData, _ = valueSSE(line)
			if err := json.Unmarshal([]byte(eventData), &event); err != nil {
				return err
			}

			if !fun(event) {
				return nil
			}
		}

	}
}
