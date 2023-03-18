package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/jeosgram/jeosgram-cli/constants"
	"github.com/jeosgram/jeosgram-cli/services"
	"github.com/jeosgram/jeosgram-cli/session"
	"github.com/jeosgram/jeosgram-cli/types"
	"golang.org/x/exp/constraints"
)

const debug = false

func isOK(res *http.Response) bool {
	return res.StatusCode == http.StatusOK
}

func isMFA(res *http.Response) bool {
	return res.StatusCode == http.StatusForbidden // 403
}

func isUnauthorized(res *http.Response) bool {
	return res.StatusCode == http.StatusUnauthorized // 401
}

func hasBadToken(body []byte) bool {
	const msg = "invalid_token"
	return bytes.Index(body, []byte(msg)) != -1
}

// -------------------

type httpValue = map[string]any

//go:generate moq -out api_test.go . JeosgramClient
type JeosgramClient interface {
	CallFunction(deviceID, funcName, funcParam string) (any, error)
	GetVariable(deviceID, varName string) (any, error)
	SignalDevice(deviceID string, signal bool)
	LoginByPassword(username, password string) (*types.Token, string, error)
	LoginByMFAOtp(mfaToken, otp string) (*types.Token, error)
	LoginByRefreshToken(refreshToken string) (*types.Token, error)
	Publish(eventName, eventData string) error
	EventStream(deviceID, eventName string, fun func(event JeosgramEvent) bool) error
}

type JeosgramAPI struct {
	token          string
	sessionService services.SessionService
}

func NewJeosgramAPI(token string, sessionService services.SessionService) *JeosgramAPI {
	return &JeosgramAPI{
		token:          "Bearer " + token,
		sessionService: sessionService,
	}
}

func (api JeosgramAPI) post(uri string, data any) (*http.Response, error) {
	body, _ := json.Marshal(data)
	return api.raw("POST", uri, body)
}

func (api JeosgramAPI) get(uri string) (*http.Response, error) {
	return api.raw("GET", uri, nil)
}

func (api JeosgramAPI) raw(method, uri string, body []byte) (*http.Response, error) {
	start := time.Now()

	url := constants.ApiURL + uri
	req, _ := http.NewRequest(
		method,
		url,
		bytes.NewReader(body),
	)
	req.Header.Set("User-Agent", constants.UserAgent)
	req.Header.Set("Authorization", api.token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json") // voy a suponer q solo se envia json
	}

	if debug {
		raw, _ := httputil.DumpRequestOut(req, true)
		fmt.Printf("\n-------\n%q\n-------\n", raw)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if debug /*&& !isOK(res)*/ {
		/*
			https://github.com/golang/go/issues/27061

			no puede obtener `transfer-encoding `
		*/
		isStream := res.Header.Get("Content-Type") == "text/event-stream"
		raw, _ := httputil.DumpResponse(res, !isStream)
		fmt.Printf("\n-------\n%q\n-------\n", raw)
	}

	if isUnauthorized(res) /*&& hasBadToken(data)*/ {
		// refrescar token y reintentar la peticion

		conf, _ := session.ReadConfig()
		token, err := api.LoginByRefreshToken(conf.RefreshToken)
		if err != nil {
			// msg requiere login
			panic(err)
		}
		api.sessionService.SaveTokens(token)

		// reintento la peticion
		return NewJeosgramAPI(token.AccessToken, api.sessionService).raw(method, uri, body)
	}

	duration := time.Since(start)
	delay := Max(constants.MinApiDelay-duration, 0)
	if debug {
		fmt.Printf("duration=%v delay=%v\n", duration, delay)
	}

	if delay > 0 {
		time.Sleep(delay)
	}

	return res, nil
}

// mover a std.Max ...
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
