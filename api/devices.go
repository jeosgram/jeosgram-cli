package api

import (
	"encoding/json"
	"fmt"
	"io"
)

func fmtUriRequestDevice(deviceID, name string) string {
	return fmt.Sprintf("/v1/devices/%s/%s", deviceID, name)
}

// -------------------------------------

func (api JeosgramAPI) CallFunction(deviceID, funcName, funcParam string) (any, error) {
	uri := fmtUriRequestDevice(deviceID, funcName)

	res, err := api.post(uri, httpValue{"arg": funcParam})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)

	if !isOK(res) {
		return nil, errorResponse(data)
	}

	var v struct {
		Value any `json:"return"`
	}
	_ = json.Unmarshal(data, &v)

	return v.Value, nil
}

func (api JeosgramAPI) GetVariable(deviceID, varName string) (any, error) {
	uri := fmtUriRequestDevice(deviceID, varName)
	res, err := api.get(uri)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)

	if !isOK(res) {
		return nil, errorResponse(data)
	}

	var v struct {
		Value any `json:"value"`
	}
	_ = json.Unmarshal(data, &v)

	return v.Value, nil
}

func (api JeosgramAPI) SignalDevice(deviceID string, signal bool) {
	// put
}
