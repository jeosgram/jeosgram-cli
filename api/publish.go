package api

import "io"

func (api JeosgramAPI) Publish(eventName, eventData string) error {
	values := httpValue{
		"name": eventName,
		"data": eventData,
		//"encode" : encode, // [hex | base16] | base32 | base64
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
