package api

import "fmt"

func fmtUriRequestDevice(deviceID, name string) string {
	return fmt.Sprintf("/v1/devices/%s/%s", deviceID, name)
}

// -------------------------------------

func (api JeosgramAPI) CallFunction(deviceID, funcName, funcParam string) {
	url := apiURL + fmtUriRequestDevice(deviceID, funcName)
	_ = url

	// post
}

func (api JeosgramAPI) GetVariable(deviceID, varName string) {
	url := apiURL + fmtUriRequestDevice(deviceID, varName)
	_ = url

	// get
}

func (api JeosgramAPI) SignalDevice(deviceID string, signal bool) {

}
