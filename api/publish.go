package api

import (
	"errors"
	"io"
	"strings"
)

func (api JeosgramAPI) Publish(eventName, eventData string) error {
	eventName = strings.TrimSpace(eventName)
	if eventName == "" {
		return errors.New("event name is required")
	}

	values := struct {
		Event  string `json:"name"`
		Data   string `json:"data,omitempty"`
		Encode string `json:"encode,omitempty"` // [hex | base16] | base32 | base64
	}{
		Event: eventName,
		Data:  eventData,
	}

	res, err := api.post("/v1/devices/events", values)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if !isOK(res) {
		data, _ := io.ReadAll(res.Body)
		return errorResponse(data)
	}

	return nil
}
